# application-01

Built GitOps using [Actions](https://docs.github.com/actions) and [Packages](https://docs.github.com/packages) of GitHub
1. When a push event occurs for all branches and tags
2. Build image
3. Push image to [Packages](https://github.com/orgs/gitops-sample/packages)
4. Edit image tag in [Helm Chart](https://github.com/gitops-sample/helm-charts)

<br/>

## When newly built
1. Token Generation
   - Scopes
     - workflow
     - write:packages
2. Create repository secrets
   - Name : ACTIONS_TOKEN
   - Secret : Token generated from 1
3. Edit `.github/workflows/gitops.yml`
   - `env.HELM_CHARTS_REPOSITORY`
   - `env.HELM_CHART_PATH`
