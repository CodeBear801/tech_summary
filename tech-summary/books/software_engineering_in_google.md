# Software engineering in Google

## Chapter 11 Testing Overview

- Why test
   + Catching bugs is only part of the motivation, an equally important reason is to **support the ability to change**
   + tests can tell you how well your entire product conforms to its intended design, and more important, when it doesn't
   + As software grows, so do test suits, face challenges like instability and slowness
   + automatic test: the best team find which to turn the collective wisdom of its members into a benefit for the entire team

- Design a test suite: the desire to reduce pain led teams to develop smaller and smaller tests
   + size: the resources that are required to run a test case, memory, processes and time
   + <img src="resources/software_engineering_in_google_C11_test_size.png" alt="software_engineering_in_google_C11_test_scope" width="600"/>
   + scope: specific code paths we are verifying
   + <img src="resources/software_engineering_in_google_C11_test_scope.png" alt="software_engineering_in_google_C11_test_scope" width="600"/>

- A better way to approach the quality of test suite is to think about the **behaviors that are tested**

- Limits of automated testing: some test need human judgment, such as search quality, video, audio.  Human explore, find problem, uncovered by probing commonly overlooked code paths or unusual responses from application, add automated test to prevent future regression.
   + Exploratory testing: which is a fundamentally creative endeavor in which someone treats the application under test as a puzzle to be broken





## Chapter 23.  Continuous Integration

- Definition: the continuous assembling and testing of entire complex and rapidly evolving ecosystem
- Google's platform: TAP
    + [CASE STUDY: The Birth Of Automated Testing At Google In 2005](https://itrevolution.com/case-study-automated-testing-google/)
- The fundamental goal of CI is to automatically catch problematic changes as early as possible
- In micro system, the changes that break an application are less likely to live inside the project's immediate codebase and more likely to be **in loosely coupled micro services on the other side of a network call**
- CI concepts
   + fast feedback loop
       * Canarying deployment
       * Experiments and feature flags are extremely powerful feedback loops
   + actionable feedback
   + automation
       * continuous build
       * continuous delivery: a continuous assembling of release candidates, followed by the promotion and testing of those candidates throughout a series of environments
  + Continuous Testing
       * presubmit's testing should be a small set: fast, reliable
       * CD-> RC -> run larger tests against the entire candidate
             
<img src="resources/software_engineering_in_google_CI_flow.png" alt="software_engineering_in_google_CI_flow" width="600"/>

- CI Challenges
- Hermetic Testing: tests run against a test environment that is entirely self-contained
     + greater determinism(stability)
     + isolation

- CI case study: Google Takeout - a data backup and download product
    + Issue 1: Continuously broken dev deploys  
        * Problem description: Takeout team development the core, many other team have their own customized deployment.  When takeout team change configuration always break the release for other team
        * The team's solution by then: Created temporary, sandbox mini-environments for each of these instances that ran on presubmit and tested that all servers were healthy on startup
        * 95% -> 50%, but not catch all -> Need end-to-end tests -> originally daily -> reused the sandboxed environments from presubmit, extending them to a new post-submit environment -> run each two hours on RC
    + Issue 2: Indecipherable test logs
        * Problem description: Takeout is the core, other products has their own plugins.  -> Takeout's end-to-end tests dumped its failures to a log -> more products, more failures -> tests always failed
        * The team's solution by then: refactor the tests(using [parameterized test runner of junit](https://github.com/junit-team/junit4/wiki/parameterized-tests)) -> clearly show tests result in UI and also attach more context information
        * Lesson learned: **Accessible, actionable feedback from CI reduces test failures and improves productivity.**
    + Issue 3: Debugging "all of google"
        * Problem description: many warning/error information are related with Google's foundation/platform
        * The team's solution by then: dedup by running CI applied to production
    + Issue 4: Keeping it green
        * Problem description: end-to-end test suites always broken and failures could not all be immediately fixed -> disable tests would make the failures too easy to forget about -> especially when rolling out a new feature for glug-ins
        * The team's solution by then: Disable failing tests by tagging them with an associated bug and filling that off the responsible team -> make tests suits green-> When rolling out a new feature, add feature flag or ID of a code change, enable a particular feature along with the output to expect both with and without the feature
        * Lesson learned: **Disabling failure tests that can't be immediately fixed is a practical approach to keeping your suite green, which gives confidence that you're aware of all test failures.**  **Automating the tests suite's maintenance, including rollout management and updating tracking bugs for fixed tests keep the suite clean and prevents technical debt** 
