package repository

import (
	"database/sql"
	"encoding/json"
	"modular-synth/internal/models"
)

type PatchRepository struct {
	DB *sql.DB
}

func (r *PatchRepository) GetPatchByID(id string) (*models.Patch, error) {
	// THIS IS WHERE YOU PLACE THE QUERY STRING
	query := `
		SELECT 
			p.id, 
			p.name,
			JSON_BUILD_OBJECT(
				'scale', p.viewport_scale,
				'offsetX', p.viewport_offset_x,
				'offsetY', p.viewport_offset_y
			) AS viewportState,
			COALESCE((
				SELECT JSON_AGG(JSON_BUILD_OBJECT(
					'id', m.id,
					'type', m.type,
					'positionX', m.position_x,
					'positionY', m.position_y,
					'parameters', m.parameters
				)) 
				FROM modules m WHERE m.patch_id = p.id
			), '[]'::json) AS modules,
			COALESCE((
				SELECT JSON_AGG(JSON_BUILD_OBJECT(
					'id', c.id,
					'sourceModule', c.source_module_id,
					'sourceHandle', c.source_handle,
					'targetModule', c.target_module_id,
					'targetHandle', c.target_handle
				)) 
				FROM connections c WHERE c.patch_id = p.id
			), '[]'::json) AS connections
		FROM patches p
		WHERE p.id = $1;
	`

	var patch models.Patch
	var vStateJSON, modulesJSON, connectionsJSON []byte

	// Execute the query
	err := r.DB.QueryRow(query, id).Scan(
		&patch.ID,
		&patch.Name,
		&vStateJSON,
		&modulesJSON,
		&connectionsJSON,
	)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSONB results into your Go structs
	json.Unmarshal(vStateJSON, &patch.ViewportState)
	json.Unmarshal(modulesJSON, &patch.Modules)
	json.Unmarshal(connectionsJSON, &patch.Connections)

	return &patch, nil
}
