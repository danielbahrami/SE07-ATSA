type api = {
  v1: {
    robot: {
      get: (id: string) => Promise<Robot>;
      getAll: () => Promise<Robot[]>;
      signal: (id: string, signal: string) => {};
    };
  };
};

const scheduling_system = process.env.SCHEDULING_SYSTEM || "localhost:22055";

console.log(scheduling_system);

const getRobot = async (id: string): Promise<Robot> => {
  const res = await fetch(`http://${scheduling_system}/api/v1/robot/${id}`);
  return await res.json();
};

const getRobots = async (): Promise<Robot[]> => {
  const res = await fetch(`http://${scheduling_system}/api/v1/robot`);
  return res.json();
};

const signalRobot = async (id: string, signal: string) => {
  await fetch(`http://${scheduling_system}/api/v1/robot/${id}/${signal}`);
};

const api: api = {
  v1: {
    robot: {
      get: getRobot,
      getAll: getRobots,
      signal: signalRobot,
    },
  },
};

export default api;
