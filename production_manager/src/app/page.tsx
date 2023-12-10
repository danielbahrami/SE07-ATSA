"use client"
import api from '../logic/api'
import {useEffect, useRef, useState} from "react";
import Card from "../components/card";
import RobotControl from '../components/robot_control';
//import {addOnMessage, sse} from "../logic/sse";

export default function Home() {

  const [robots, setRobots] = useState<Robot[]>([])
  const [notification, setNotification] = useState<string|undefined>(undefined)

 /* addOnMessage((data) => {
    setNotification(data);
  });*/

  useEffect(() => {
    const sse = new EventSource("http://127.0.0.1:5000/sse");
    sse.onmessage = e => {
      console.log(e.data)
      if (e.data != "") {
        setNotification(e.data);
      }
    }
    api.v1.robot.getAll().then(rs => {
      setRobots(rs.sort((a,b) => comp(a,b)))
    })
  }, [])

  const comp = (a: Robot, b: Robot): number => {
    console.log(a.ID, b.ID);
    
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
          <div className="state" style={{display: 'flex'}}>
            <span>State: </span><div className="state-indicator" style={{background: stateColor(r.State), width: "1rem", height: "1rem", borderRadius: "50%"}}/>
          </div> <RobotControl robot={r}/>
        </div>}></Card>)}
    </main>
  )
}
