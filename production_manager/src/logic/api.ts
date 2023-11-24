type api = {
  v1: {
    schedules: {
      get: (id: string) => Promise<Schedule>;
      getAll: () => Promise<Schedule[]>;
      post: (schedule: Schedule) => any;
      patch: () => any;
      delete: () => any;
    };
    productionLine: {
      get: (id: string) => Promise<ProductionLine>;
      getAll: () => Promise<ProductionLine[]>;
      post: (productionLine: ProductionLine) => any;
      patch: () => any;
      delete: () => any;
    };
  };
};

const GetSchedule = async (id: string): Promise<Schedule> => {
  const res = await fetch(`http://127.0.0.1:22055/api/v1/schedule/${id}`, {
    method: "GET",
  });
  return await res.json();
};

const GetAllSchedules = async (): Promise<Schedule[]> => {
  const res = await fetch("http://127.0.0.1:22055/api/v1/schedule/", {
    method: "GET",
  });
  return await res.json();
};

const PostSchedule = async (schedule: Schedule) => {
  console.log("body", JSON.stringify(schedule));
  const res = await fetch(`http://127.0.0.1:22055/api/v1/schedule`, {
    method: "POST",
    headers: {
      "Content-Type": "Application/json",
    },
    body: JSON.stringify(schedule),
  });
  return await res.json();
};

const PatchSchedule = () => {};

const DeleteSchedule = () => {};

const GetProductionLine = async (id: string): Promise<ProductionLine> => {
  const res = await fetch(
    `http://127.0.0.1:22055/api/v1/production_line/${id}`,
    {
      method: "GET",
    }
  );
  return await res.json();
};

const GetAllProductionLines = async (): Promise<ProductionLine[]> => {
  const res = await fetch("http://127.0.0.1:22055/api/v1/production_line", {
    method: "GET",
  });
  return await res.json();
};

const PostProductionLine = async (productionLine: ProductionLine) => {
  console.log("body", JSON.stringify(productionLine));
  const res = await fetch(`http://127.0.0.1:22055/api/v1/production_line`, {
    method: "POST",
    headers: {
      "Content-Type": "Application/json",
    },
    body: JSON.stringify(productionLine),
  });
  return await res.json();
};

const PatchProductionLine = () => {};

const DeleteProductionLine = () => {};

const api: api = {
  v1: {
    schedules: {
      get: GetSchedule,
      getAll: GetAllSchedules,
      post: PostSchedule,
      patch: PatchSchedule,
      delete: DeleteSchedule,
    },
    productionLine: {
      get: GetProductionLine,
      getAll: GetAllProductionLines,
      post: PostProductionLine,
      patch: PatchProductionLine,
      delete: DeleteProductionLine,
    },
  },
};

export default api;
