import "./App.css";

import { Button, Form } from "react-bootstrap";
import { useState } from "react";
import Basketball from "./Basketball";

function App() {
  const [gameNames, setGameNames] = useState("");
  const [showGame, setShowGame] = useState(false);

  const handleChange = (e) => {
    setGameNames(e.target.value);
  };

  const handleMatches = async (e) => {
    e.preventDefault();
    const gameNamesBody = gameNames.replaceAll(" ", "").split(",");
    await fetch("/games", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        redirect: "follow",
      },

      body: JSON.stringify(gameNamesBody),
    }).then((response) => {
      if (response) {
        setShowGame(true);
      }
    });
  };

  return (
    <div className="container">
      {!showGame && (
        <Form onSubmit={handleMatches}>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>mathc1, match2, ...</Form.Label>
            <Form.Control
              type="text"
              value={gameNames}
              onChange={handleChange}
            />
          </Form.Group>
          <Button variant="warning" type="submit">
            Start Matches
          </Button>
        </Form>
      )}
      {showGame ? <Basketball /> : null}
    </div>
  );
}

export default App;
