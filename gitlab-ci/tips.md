# GitLab CI Tips

## YAML File

### Hide Job
Prepent jobs with `.` to make then hidden jobs, and they are skipped by GitLab CI/CD.
Ref: https://docs.gitlab.com/ee/ci/jobs/index.html#hide-jobs
eg.
```yaml
.hidden_job:
  script:
    - run test
```

### List GitLab-CI Variables
Ref: https://docs.gitlab.com/ee/ci/variables/index.html#list-all-environment-variables
```yaml
build_name:
  script:
    - export
```