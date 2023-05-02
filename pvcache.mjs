import { createServer } from "http";
const fallBack = JSON.stringify({ Body: { Data: { Site: { P_PV: null } } } });

const server = createServer(async function (_, res) {
  console.log(`request`);
  try {
    const realtime = await fetch(
      "http://pv.fritz.box/solar_api/v1/GetPowerFlowRealtimeData.fcgi"
    );
    for (const header of realtime.headers.entries()) {
      res.setHeader(header[0], header[1]);
    }
    res.writeHead(200);
    res.end(await realtime.text());
  } catch {
    res.setHeader("content-type", "application/json");
    res.setHeader("cache-control", "no-cache, no-store, must-revalidate");
    res.setHeader("pragma", "no-cache");
    res.writeHead(200);
    res.end(fallBack);
  }
});
server.listen(8000, "0.0.0.0", () => {
  console.log(`Server is running`);
});
