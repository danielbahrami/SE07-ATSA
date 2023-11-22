"use client"
import api from '../logic/api'
import {useEffect, useRef, useState} from "react";
import Card from "../components/card";

export default function Home() {

  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [productionLines, setProductionLines] = useState<ProductionLine[]>([]);

  const scheduleNameField = useRef<HTMLInputElement>();
  const productionLineNameField = useRef<HTMLInputElement>();

  useEffect(() => {
    api.v1.schedules.getAll().then(schs => {
      console.log(schs)
      setSchedules(schs)
    })
    api.v1.productionLine.getAll().then(plns => {
      console.log(plns)
      setProductionLines(plns)
    })
  }, []);


  const createSchedule = (event) => {
    event.preventDefault();
    if (!scheduleNameField.current) {
      return;
    }
    const schedule: Schedule = {
      ID: `sch_${schedules.length+1}`,
      Name: scheduleNameField.current!.value,
      Description: "DESCRIPTION",
      ProductionLineId: "prod_01"
    }
    console.log(schedule);
    api.v1.schedules.post(schedule).then(r => {
      console.log(r)
      schedules.push(schedule)
      setSchedules([...schedules])
      console.log(schedules)
    });
  }

  const createProductionLine = (event) => {
    event.preventDefault();
    if (!productionLineNameField.current) {
      return;
    }
    const productionLine: ProductionLine = {
      ID: `pln_${productionLines.length+1}`,
      Name: productionLineNameField.current!.value,
      Status: "Standby",
    }
    console.log(productionLine);
    api.v1.productionLine.post(productionLine).then(r => {
      console.log(r);
      productionLines.push(productionLine);
      setProductionLines([...productionLines]);
      console.log(productionLines)
    })
  }

  return (
    <main>
      {schedules.map((sch) => <Card key={`key_${sch.ID}`} title={"Schedule"} body={<div>{sch.Name} ({sch.ID})</div>}/>)}

      <form onSubmit={createSchedule}>
        <input type="text" name="schedule_id" id="schedule_id" ref={scheduleNameField}/>
        <input type="submit" value="Create Sch"/>
      </form>

      {productionLines.map((pln) => <Card key={`key_${pln.ID}`} title={"ProductionLine"} body={<div>{pln.Name} Status: {pln.Status}</div>}/>)}

      <form onSubmit={createProductionLine}>
        <input type="text" name="production_line_name" id="production_line_id" ref={productionLineNameField}/>
        <input type="submit" value="Create Pln"/>
      </form>
    </main>
  )
}
