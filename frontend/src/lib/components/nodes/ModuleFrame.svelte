<script lang="ts">

    let { title, children } = $props();
</script>

<div class="module-panel rail-border">
    <div class="panel-header">
        <div class="screw"></div>
        <span>{title}</span>
        <div class="screw"></div>
    </div>

    <div class="panel-content">
    {@render children()}
  </div>
</div>

<style>
  .module-panel {
    width: 120px;
    height: 300px;
    background: #2a2a2a;
    border: 1px solid #333;
    box-shadow: inset 0 0 10px rgba(0,0,0,0.5), 2px 2px 5px rgba(0,0,0,0.3);
    border-radius: 2px;
    display: flex;
    flex-direction: column;
    padding: 10px;
  }

  .screw {
    width: 8px;
    height: 8px;
    background: #111;
    border: 1px solid #444;
    border-radius: 50%;
    background-image: radial-gradient(circle, #555 1px, transparent 2px);
  }
</style>