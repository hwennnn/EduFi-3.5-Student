# EduFi-3.5-Student

## Folder Structure

|       Codebase       |            Description             |
| :------------------: | :--------------------------------: |
| [frontend](frontend) |       React Next.js Frontend       |
|  [backend](backend)  |           EduFi Backend            |
|   [server](server)   | General-Purpose API Backend Server |

## Usage

### Development

```bash
docker-compose pull
docker-compose up
```

**The SQL setup script will be executed automatically at build time.**

Finally, open http://localhost:9210 for the frontend client.

### Run in Production Mode

```bash
docker-compose -f docker-compose.prod.yml pull
docker-compose -f docker-compose.prod.yml up
```

## Architecture Diagram

![Architecture Diagram](docs/architecture_diagram.png)

## API Documentation

Kindly refer to [this](docs/README.md) for more details on the EduFi Student API server.

## **Credits**

<table>
  <tr>
    <td align="center"><a href="https://github.com/hwennnn"><img src="https://avatars3.githubusercontent.com/u/54523581?s=460&u=a649d3ed6c70ffe2fa69f37c0870415668149113&v=4" width="100px;" alt=""/><br /><sub><b>Wai Hou Man <br> (S10197636F) </b></sub></a><br />
    </td>
  </tr>
</table>
