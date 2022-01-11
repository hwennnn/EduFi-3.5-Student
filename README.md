# EduFi-3.5-Student

## Folder Structure

|   Codebase   | Description |
| :----------: | :---------: |
| [REST](REST) |  REST API   |

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
