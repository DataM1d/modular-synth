<script lang="ts">
  import { Handle, Position } from '@xyflow/svelte';
  import ModuleFrame from './ModuleFrame.svelte';
  import Knob from '../ui/Knob.svelte';

  let {  } = $props();
</script>

<ModuleFrame title="OSC-1">
  <div class="controls">
    <label for="freq">FREQ</label>
    <Knob />
    <label for="level">LEVEL</label>
    <Knob />
  </div>

  <Handle type="source" position={Position.Right} id="out" />
  <Handle type="target" position={Position.Left} id="fm" />
</ModuleFrame>

<style>
  .controls {
    display: flex;
    flex-direction: column;
    gap: 15px;
    align-items: center;
  }
  label { font-size: 8px; color: #666; margin-top: 5px; }
</style>