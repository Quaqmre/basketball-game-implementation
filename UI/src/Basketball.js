import React, { useEffect, useState } from "react";
import { Table, Card, ListGroup, Button } from "react-bootstrap";

function Basketball() {
  const [gameDatas, setGameDatas] = useState([]);

  const handleData = async () => {
    await fetch("/games", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((response) => {
      if (response) {
        response.text().then((data) => {
          setGameDatas(JSON.parse(data));
        });
      }
    });
  };

  useEffect(() => {
    handleData();
    setInterval(() => {
      handleData();
    }, 5000);
    return () => clearInterval();
  }, []);

  console.log(gameDatas);
  return (
    <div>
      {gameDatas.map((game, index) => (
        <div key={index}>
          <Card>
            <Card.Body>
              <Card.Title>{`Game Name: ${game.Name}`}</Card.Title>
              <Card.Title>{`Match Statistics`}</Card.Title>
              {game.TopScorer && game.TopAssistPlayer && (
                <ListGroup as="ol">
                  <ListGroup.Item
                    as="li"
                    className="d-flex justify-content-start align-items-start"
                  >
                    <div className="ms-2 ">
                      <div className="fw-bold">Top Assist Player</div>
                      {`Name: ${game.TopAssistPlayer.Name}`}
                      <br />
                      {`Score: ${game.TopAssistPlayer.Assist}`}
                    </div>
                  </ListGroup.Item>
                  <ListGroup.Item
                    as="li"
                    className="d-flex justify-content-start align-items-start"
                  >
                    <div className="ms-2 ">
                      <div className="fw-bold">Top Scorer</div>
                      {`Name: ${game?.TopScorer?.Name}`}
                      <br />
                      {`Score: ${game?.TopScorer?.Score}`}
                    </div>
                  </ListGroup.Item>{" "}
                </ListGroup>
              )}
              {game.Teams.map((team, index) => (
                <Card key={index}>
                  <Card.Body>
                    <Card.Title>{`Team Name: ${team.Name}`}</Card.Title>
                    <Card.Title>{`Team Point: ${team.Point}`}</Card.Title>
                    <Card.Title>{`Team Attack Count: ${team.AttackCount}`}</Card.Title>
                    PLAYERS
                    <Table striped bordered hover>
                      <thead>
                        <tr>
                          <th>#</th>
                          <th>Player Name</th>
                          <th>Score</th>
                          <th>Assist</th>
                          <th>Ability</th>
                        </tr>
                      </thead>
                      <tbody>
                        {team.Players.map((player, index) => (
                          <tr key={index}>
                            <td>{index + 1}</td>
                            <td>{player.Name}</td>
                            <td>{player.Score}</td>
                            <td>{player.Assist}</td>
                            <td>{player.Ability}</td>
                          </tr>
                        ))}
                      </tbody>
                    </Table>
                  </Card.Body>
                </Card>
              ))}
            </Card.Body>
          </Card>
          <hr className="dashed" />
          <hr className="dashed" />
        </div>
      ))}
    </div>
  );
}

export default Basketball;
