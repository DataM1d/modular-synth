-- Core Patch container
CREATE TABLE patches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    viewport_scale FLOAT DEFAULT 1.0,
    viewport_offset_x FLOAT DEFAULT 0.0,
    viewport_offset_y FLOAT DEFAULT 0.0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Modules associated with a specific Patch
CREATE TABLE modules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patch_id UUID REFERENCES patches(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    position_x FLOAT NOT NULL,
    position_y FLOAT NOT NULL,
    parameters JSONB DEFAULT '{}'::jsonb
);

-- Connections linking module handles
CREATE TABLE connections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patch_id UUID REFERENCES patches(id) ON DELETE CASCADE,
    source_module_id UUID REFERENCES modules(id) ON DELETE CASCADE,
    source_handle TEXT NOT NULL,
    target_module_id UUID REFERENCES modules(id) ON DELETE CASCADE,
    target_handle TEXT NOT NULL,
    -- Ensure connection is within the same patch
    CONSTRAINT fk_patch FOREIGN KEY (patch_id) REFERENCES patches(id) ON DELETE CASCADE
);

-- Indexing for performance
CREATE INDEX idx_modules_patch_id ON modules(patch_id);
CREATE INDEX idx_connections_patch_id ON connections(patch_id);