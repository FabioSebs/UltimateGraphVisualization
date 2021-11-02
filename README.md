# :fire: :dragon: Ultimate Graph Visualization :fire: :dragon:

---

## How to Run (2 Ways) :runner:

### Docker :whale2:

> This avoids having to install Go and projetc dependencies on your local machine and assumes you have docker installed

```bash
$docker pull fabrzy/graph
$docker images
$docker run {imageid}
```

The **_docker images_** command should show you the image id that you have pulled from dockerhub repo. The **_image id_** is used for the **_docker run_** command.

### Git Clone :git:

> this assumes you have Go installed.

```bash
$git clone https://github.com/FabioSebs/UltimateGraphVisualization.git
$cd UltimateGraphVisualization
$go get -u
$go run .
```
