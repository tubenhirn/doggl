const commitAnalyzerOptions = {};
const releaseNotesGeneratorOptions = {};

module.exports = {
  branches: ["main"],
  plugins: [
    ["@semantic-release/commit-analyzer", commitAnalyzerOptions],
    ["@semantic-release/release-notes-generator", releaseNotesGeneratorOptions],
    [
      "@semantic-release/changelog",
      {
        changelogFile: "CHANGELOG.md",
      },
    ],
    [
      "semantic-release-replace-plugin",
      {
        replacements: [
          {
            files: ["version"],
            from: "v.*",
            to: "v${nextRelease.version}",
            results: [
              {
                file: "version",
                hasChanged: true,
                numMatches: 1,
                numReplacements: 1,
              },
            ],
            countMatches: true,
          },
        ],
      },
    ],
    [
      "@semantic-release/git",
      {
        assets: ["version", "CHANGELOG.md"],
      },
    ],
    ["@semantic-release/github"],
  ],
};
