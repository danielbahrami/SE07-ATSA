import { FunctionComponent } from "react";
import api from "../logic/api";

interface RobotControlProps {
  robot: Robot;
}

const RobotControl: FunctionComponent<RobotControlProps> = (
  props: RobotControlProps
) => {
  const startRobot = () => {
    api.v1.robot.signal(props.robot.ID, "start");
  };

  const stopRobot = () => {
    api.v1.robot.signal(props.robot.ID, "halt");
  };

  const failRobot = () => {
    api.v1.robot.signal(props.robot.ID, "xd");
  };

  return (
    <div>
      <button onClick={startRobot}>Start</button>
      <button onClick={stopRobot}>Stop</button>
      <button onClick={failRobot}>Fail</button>
    </div>
  );
};

export default RobotControl;
