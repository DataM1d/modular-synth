The Evolution of the Project

This project began as an exploration of 3D audio sequencing under the name Orbita Clone, leveraging Three.js, React, and TypeScript. While the initial build provided a great visual playground, the heavy abstraction layers and the overhead of 3D rendering eventually created performance bottlenecks that hindered the fluidity required for real time audio synthesis.

Drawing inspiration from the tactile and logical flow of VCV Rack, I decided to pivot. I rebuilt the architecture from the ground up to prioritize performance and responsiveness:

Front end Shift to Svelte: Moved from React to Svelte to eliminate the virtual DOM overhead. This change provides surgical UI updates, allowing the interface to remain snappy even as the number of modules and patches increases.

Back end Migration to Go: Replaced the legacy stack with a robust Go backend. This ensures the system can handle concurrent connections and complex backend logic with the efficiency required for a modern, scalable application.

This transition reflects my commitment to selecting the right tools for the specific demands of an application prioritizing raw performance and clean architecture over convenience.