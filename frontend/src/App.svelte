<script lang="ts">
  import { SvelteFlow, SvelteFlowProvider, Controls } from '@xyflow/svelte';
  import type { Node, Edge } from '@xyflow/svelte';
  import '@xyflow/svelte/dist/style.css';
  import OscillatorNode from './lib/components/nodes/OscillatorNode.svelte';
  import CustomBackground from './lib/components/CustomBackground.svelte';

  const nodeTypes = { oscillator: OscillatorNode };

  let nodes: Node[] = $state([
    { id: '1', type: 'oscillator', position: { x: 100, y: 100 }, data: {} }
  ]);
  let edges: Edge[] = $state([]);
</script>

<div class="rack-workspace">
  <SvelteFlowProvider>
    <div class="custom-bg-container">
      <CustomBackground />
    </div>
    
    <SvelteFlow 
      {nodes} 
      {edges} 
      {nodeTypes}
    >
      <Controls />
    </SvelteFlow>
  </SvelteFlowProvider>
</div>

<style>
  .rack-workspace {
    position: relative;
    width: 100vw;
    height: 100vh;
    background-color: #0f061b;
    overflow: hidden;
  }

  .custom-bg-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
    pointer-events: none;
  }

  :global(.svelte-flow) {
    background: transparent !important;
  }
</style>