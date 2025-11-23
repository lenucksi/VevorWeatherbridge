Summary - GitHub Release Improvements Complete ✅
What Was Fixed
You reported that GitHub releases were very blank with just "Release notes for version 0.1.7" - missing changelog, commits, contributors, and Docker image info.
Changes Made
Updated v0.1.7 Release - Now includes:
Full changelog from CHANGELOG.md
Critical fix explanation (MQTT Discovery)
List of all 11 sensors with emojis
Installation instructions (HA Addon + Docker)
Docker image pull commands for all architectures
Contributors list
Full commit link
Updated v0.1.6 Release - Now includes:
Bashio migration details
Why Bashio section with benefits
Installation instructions
Contributors and commit info
Created Automated Release Workflow - .github/workflows/release.yml
Triggers automatically on version tags (v*.*.*)
Extracts changelog section from CHANGELOG.md
Generates contributor list from git history
Lists all commits since previous release
Creates comparison link to previous version
Includes Docker installation commands
Fully automated - no manual work needed
How It Works
For Future Releases:
Update version in config.yaml
Add entry to CHANGELOG.md
Commit changes
Create and push git tag: git tag v0.1.8 && git push origin v0.1.8
GitHub Actions automatically creates comprehensive release notes
Current Releases:
View updated v0.1.7: https://github.com/lenucksi/VevorWeatherbridge/releases/tag/v0.1.7
View updated v0.1.6: https://github.com/lenucksi/VevorWeatherbridge/releases/tag/v0.1.6
All releases now have complete information including changelog, contributors, commits, installation instructions, and Docker image references.