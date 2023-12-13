"use client"
import api from '../logic/api'
import { useEffect, useRef, useState } from "react";
import Card from "../components/card";
import RobotControl from '../components/robot_control';

export default function Home() {

  const updateRobot = (rs: Robot[], data: string) => {
    console.log(rs);
    const lines = data.split(",");
    const id = lines[0].split("=")[1];
    const state = lines[1].split("=")[1];
    console.log("id", id, "state", state)
    const updatedRobots = rs.map(r => {
      console.log(r.ID)
      if (r.ID === id) {
        r.State = state;
      }
      return r
    })
    setRobots(updatedRobots);
    console.log("new", updatedRobots);
  }
  const nofification_system = process.env.NOTIFICATION_SYSTEM || "localhost:55022";

  const [robots, setRobots] = useState<Robot[]>([])
  const [notification, setNotification] = useState<string | undefined>("Welcome")

  useEffect(() => {
    
    api.v1.robot.getAll().then(rs => {
      let arr = rs;
      setRobots(rs.sort((a,b) => comp(a,b)))

      const sse = new EventSource(`http://${nofification_system}/sse`);
      sse.onmessage = e => {
        console.log(e.data)
        if (e.data != "") {
          setNotification(e.data);
          updateRobot(arr, e.data)
        }
      }
      let count = 1;
      sse.onerror = e => {
        setNotification(`Could not connect to notification system, retrying ... Attempt [${count++}/3]`)
        if (count > 4) {
          sse.close();
          setNotification("Notification system is unavailable")
        }
      }

    });
        
  }, [])

  const comp = (a: Robot, b: Robot): number => {
    return Number(a.ID.split("_")[1]) - Number(b.ID.split("_")[1])
  }

  const stateColor = (state: string): string => {
    switch (state) {
      case "STARTING":
        return "cyan"
      case "RUNNING":
        return "green"
      case "STOPPING":
        return "orange"
      case "IDLE":
        return "yellow"
      default:
        return "red"
    }
  }

  return (
    <main>
      {notification &&
        <div>{notification}</div>}
      <h1>Production manager</h1>
      {robots.map(r =>
        <Card key={`key_${r.ID}`} title="Robot" body={
          <div>
            ID: {r.ID}
            <div className="state" style={{ display: 'flex' }}>
              <span>State: </span><div className="state-indicator" style={{ background: stateColor(r.State), width: "1rem", height: "1rem", borderRadius: "50%" }} />
            </div> <RobotControl robot={r} />
          </div>}></Card>)}
    </main>
  )
}