## Description

<!-- Brief summary of the changes and their purpose. -->

## Type of Change

- [ ] Bug fix (non-breaking change that fixes an issue)
- [ ] New feature (non-breaking change that adds functionality)
- [ ] Refactor (no functional changes, no API changes)
- [ ] Test improvement (new tests, renamed tests, coverage increase)
- [ ] Documentation (README, comments, memory files)
- [ ] CI/CD (workflow, pipeline, automation changes)
- [ ] Breaking change (fix or feature that would cause existing functionality to change)

## Checklist

- [ ] Code compiles without errors (`./run.ps1 PC`)
- [ ] All tests pass (`./run.ps1 TC`)
- [ ] No new `// Deprecated:` markers without migration plan
- [ ] Version bumped (S-015 discipline)
- [ ] `.release` folder not modified
- [ ] Test naming follows `Test_{TypeOrFunc}_{Scenario}` convention
- [ ] Assertions use `args.Map` + `ShouldBeEqual` (not raw `t.Fatal`/`t.Error`)

## Linked Issues

<!-- Reference related issues: Fixes #123, Relates to #456 -->

## Additional Notes

<!-- Any context reviewers should know: trade-offs, follow-up work, etc. -->
