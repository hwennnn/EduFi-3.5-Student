# EduFi-3.5-Student

## Folder Structure

|      Codebase      |            Description             |
| :----------------: | :--------------------------------: |
| [backend](backend) |           EduFi Backend            |
|  [server](server)  | General-Purpose API Backend Server |

## Usage

### On Docker-Compose

```bash
docker-compose pull
docker-compose up
```

### Initialise the MySQL database with setup script

```bash
docker exec -it edufi_student_db bash
mysql -uroot -p Edufi_Student < app/setup.sql
```

Finally, open http://localhost:80 for the frontend client.
