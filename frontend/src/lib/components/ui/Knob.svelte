<script lang="ts">
    let value = $state(0.5);

    function handleMouseMove(e: MouseEvent) {
        const change = -e.movementY * 0.005;
        value = Math.min(Math.max(value + change, 0), 1);
    }
</script>

<div class="knob" 
     role="slider"
     tabindex="0"
     aria-valuemin="0"
     aria-valuemax="1"
     aria-valuenow={value}
     onmousedown={() => window.addEventListener('mousemove', handleMouseMove, { once: false })}
     onmouseup={() => window.removeEventListener('mousemove', handleMouseMove)}>
  <div class="dial" style="transform: rotate({value * 270 - 135}deg)"></div>
</div>

<style>
  .knob { width: 30px; height: 30px; border-radius: 50%; background: #111; border: 2px solid #555; cursor: ns-resize; position: relative; }
  .dial { position: absolute; width: 2px; height: 10px; background: #0f0; top: 2px; left: 14px; transform-origin: 50% 13px; }
</style>