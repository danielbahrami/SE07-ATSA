import React, { FunctionComponent } from "react";
import styles from "./card.css";

interface OwnProps {
  title: string;
  body: any;
}

type Props = OwnProps;

const Card: FunctionComponent<Props> = (props) => {
  return (
    <div className="card" style={styles.card}>
      <div>{props.title}</div>
      <div>{props.body}</div>
    </div>
  );
};

export default Card;
