# Score API

We can think this project as a micro service for calculating and storing player's scores.

## Endpoints

- <code>GET</code> /api/scores returns all player scores
- <code>GET</code> /api/scores/:id returns all scores of a player by id
- <code>POST</code> /api/scores creates a new score

  ```json
   { // request body
     "match_id": uint,
     "player_id": uint,
     "kill_count": uint,
     "survival_time": uint (ms)
   }
  ```

## TODO
- docker compose database connection
- auth system
- common response
