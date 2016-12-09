Contributing guidelines
=======================

### Git workflow

* We use git-flow - create a feature branch from `develop`, e.g. `feature/new-feature`
* Pull requests must contain a succinct, clear summary of what the user need is driving this feature change
* Ensure your branch contains logical atomic commits before sending a pull request - follow the [alphagov Git styleguide](https://github.com/alphagov/styleguides/blob/master/git.md)
* You may rebase your branch after feedback if it's to include relevant updates from the develop branch. We prefer a rebase here to a merge commit as we prefer a clean and straight history on develop with discrete merge commits for features

### Adding a new model

See [the wiki page](https://github.com/ONSdigital/dp-frontend-models/wiki/Adding-a-new-model-package-and-updating-dependent-projects)
for detailed instructions on how to add a new model and update dependent projects. In short: add your new model package
and commit locally, then run `govendor add github.com/ONSdigital/dp-frontend-models/model/mynewmodel` in the dependent
project to add it to the vendor dependencies.