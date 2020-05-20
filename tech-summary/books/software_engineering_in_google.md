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
   + scope: how much code a test is intended to validate
   + <img src="resources/software_engineering_in_google_C11_test_scope.png" alt="software_engineering_in_google_C11_test_scope" width="600"/>

- A better way to approach the quality of test suite is to think about the **behaviors that are tested**

- Limits of automated testing: some test need human judgment, such as search quality, video, audio.  Human explore, find problem, uncovered by probing commonly overlooked code paths or unusual responses from application, add automated test to prevent future regression.
   + Exploratory testing: which is a fundamentally creative endeavor in which someone treats the application under test as a puzzle to be broken


## Chapter 12 Unit Testing
- Maintainable tests: after writing them, engineers don't need to think about them again until they fail, and those failures indicate real bugs with clear causes

### Avoid brittle tests

- Brittle tests: the one that fails in the face of an unrelated change to production code that does not introduce any real bugs.
   + Only breaking changes in a system's behavior should require going back to change its tests, and in such situations, the cost of updating those tests tend to be small relative to the cost of updating all of the system's users


- [**Test via public API**](https://testing.googleblog.com/2015/01/testing-on-toilet-prefer-testing-public.html)
   + public API means **explicit contracts**, they exposed by that unit to third parties outside of the team that owns the code
   + If a method or class exists only to support one or two other classes, it probably shouldn't be considered its own unit
   + If a package or class is designed to be accessible by anyone without having to consult with its owners, must need a unit test
   + If a package or class can be accessed only by the people who own it, but it is designed to provide a general piece of functionality useful in a range of contexts("support library"), better also consider as a unit and test directly

```java
public void processTransaction(Transaction transaction) {
	if (isValid(transaction)) {
		saveToDatabase(transaction);
	}
}

private boolean isValid(Transaction t) {
	// ...
}

private void saveToDatabase(Transaction t) {
	// ...
}

public void setAccountBalance(String accountName, int balance) {
	// Write the balance to the database
}

public volid getAccountBalance(String accountName) {
	// Read transactions from database
}

```

```java
@Test
public void shouldTransferFunds() {
    processor.setAccountBalance("me", 150);
    processor.setAccountBalance("you", 20);

    processor.processTransaction(newTransaction()
        .setSender("me")
        .setRecipient("you")
        .setAmount(100));

    assertThat(processor.getAccountBalance("me")).isEqualTo(50);
    assertThat(processor.getAccountBalance("you")).isEqualTo(120);
}

@Test
public void shouldNotPerformInvalidTransactions() {
    processor.setAccountBalance("me", 50);
    processor.setAccountBalance("you", 20);

        processor.processTransaction(newTransaction()
        .setSender("me")
        .setRecipient("you")
        .setAmount(100));

    assertThat(processor.getAccountBalance("me")).isEqualTo(50);
    assertThat(processor.getAccountBalance("you")).isEqualTo(20);
}

```
- [**Test State, not interactions**](https://testing.googleblog.com/2013/03/testing-on-toilet-testing-state-vs.html)
   + With state testing, you observe the system itself to see what it looks like after invoking with it.
   + With interaction testing, you instead check that the system took an expected sequence of actions on its collaborators in response to invoking it
   + interaction tests check **how** a system arrived at its result, whereas usually you should care only **what** the result is

```java
public void testSortNumbers() {
  NumberSorter numberSorter = new NumberSorter(quicksort, bubbleSort);
  // Verify that the returned list is sorted. It doesn't matter which sorting
  // algorithm is used, as long as the right result is returned.
  assertEquals(
      new ArrayList(1, 2, 3),
      numberSorter.sortNumbers(new ArrayList(3, 1, 2)));
}
```

### Writing clear tests

- A clear test is one whose purpose for existing and reason for falling is immediately clear to the engineer diagnosing a failure
   + Unclear tests always result in be dropped off, introducing a subtle hole in test coverage

- [**Make your tests complete and concise**](https://testing.googleblog.com/2014/03/testing-on-toilet-what-makes-good-test.html)
   + test is complete when its body contains all of the information a reader needs in order to understand how it arrives at its result
   + test is concise when it contains no other distraction or irrelevant information
   + **a resilient test doesn't have to change unless the purpose or behavior of the class being tested changes.**
   + a test's body should contain all of the information needed to understand it without containing any irrelevant or distracting information

```java
// incomplete and cluttered test
@Test public void shouldPerformAddition() {
  Calculator calculator = new Calculator(new RoundingStrategy(), 
      "unused", ENABLE_COSIN_FEATURE, 0.01, calculusEngine, false);
  int result = calculator.doComputation(makeTestComputation());
  assertEquals(5, result); // Where did this number come from?
}

// Lots of distracting information is being passed to the constructor, and the important parts are hidden off in a helper method.

// complete, concise test
@Test public void shouldPerformAddition() {
  Calculator calculator = newCalculator();
  int result = calculator.doComputation(makeAdditionComputation(2, 3));
  assertEquals(5, result);
}

```
- [**Test behaviors, not methods**](https://testing.googleblog.com/2014/04/testing-on-toilet-test-behaviors-not.html)
```java
// a transaction snippet
@Test public void testProcessTransaction() {
  User user = newUserWithBalance(LOW_BALANCE_THRESHOLD.plus(dollars(2));
  transactionProcessor.processTransaction(
      user,
      new Transaction("Pile of Beanie Babies", dollars(3)));
  assertContains("You bought a Pile of Beanie Babies", ui.getText());
  assertEquals(1, user.getEmails().size());
  assertEquals("Your balance is low", user.getEmails().get(0).getSubject());
}

// a method driven test
@Test public void testDisplayTransactionResults() {
    transactionProcessor.displayTransactionResults(
        newUserWithBalance(LOW_BALANCE_THRESHOLD.plus(dollars(2))),
        new Transaction("some item", dollars(3))
    );
    assertThat(ui.getText()).contains("You bought a some item");
    assertThat(ui.getTest()).contains("Your balance is low")
}

// a behavior driven test
@Test public void testProcessTransaction_displaysNotification() {
  transactionProcessor.processTransaction(
      new User(), new Transaction("Pile of Beanie Babies"));
  assertContains("You bought a Pile of Beanie Babies", ui.getText());
}
@Test public void testProcessTransaction_sendsEmailWhenBalanceIsLow() {
  User user = newUserWithBalance(LOW_BALANCE_THRESHOLD.plus(dollars(2));
  transactionProcessor.processTransaction(
      user,
      new Transaction(dollars(3)));
  assertEquals(1, user.getEmails().size());
  assertEquals("Your balance is low", user.getEmails().get(0).getSubject());
}

```
- rather than writing a test for each method, write a test for each behavior.  
- A behavior is any gurantee that a system makes about how it will respond to a series of inputs while in a particular state
- [**Given, When, Then**](https://martinfowler.com/bliki/GivenWhenThen.html)
   + **given** defines how the system is set up
   + **when** defines the action to be taken on the system
   + **then** validates the result
   + [Cucumber](https://cucumber.io/), [spock](http://spockframework.org/)
- Split tests to [**keep them more focused**](https://testing.googleblog.com/2018/06/testing-on-toilet-keep-tests-focused.html)
- Behavior driven tests tend to be clearer
   + They read more like natural language, easy to be understood
   + More clearly express [cause and effect](https://testing.googleblog.com/2017/01/testing-on-toilet-keep-cause-and-effect.html) because each test is more limited in scope
   + the fact that each test is short and descriptive makes it easier to see what functionality is already tested and encourages engineer to add new streamlined test methods instead of piling onto existing methods

```java
@Test public void shouldTimeOutConnections() {
    // Given two users
    User user1 = newUser();
    User user2 = newUser();

    // And an empty connection pool with a 10-minute timeout
    Pool pool = newPool(Duration.minutes(10));
    
    // When connecting both users to the pool
    pool.connect(user1);
    pool.connect(user2);
    
    // Then the pool should have two connections
    assertThat(pool.getConnections()).hasSize(2);

    // When waiting for 20 minutes
    clock.advance(Duration.minutes(20));

    // Then the pool should have no connections
    assertThat(pool.getConnections()).isEmpty();

    // And each user should be disconnected
    assertThat(user1.isConnected()).isFalse();
    assertThat(user2.isConnected()).isFalse();
}
```

- [**Define clear test name**](https://testing.googleblog.com/2014/10/testing-on-toilet-writing-descriptive.html).  A good name describes both the actions that are being taken on a system and the expected outcome.
   + Test names will sometimes include additional information like the state of the system or its environment before taking action on it
   + A good trick if you're stuck is to try starting the test name with the word **should**.  When taken with the name of the class being tested, this naming scheme allows the test name to be read as a sentence.

```
shouldNotAllowWithdrawalsWhenBalanceIsEmpty
shouldLockOutUserAfterThreeInvalidLoginAttempts
```

- [**Don't put logic into tests**](https://testing.googleblog.com/2014/07/testing-on-toilet-dont-put-logic-in.html)
   + good test: a test is doing the correct thing just from glancing at it
   + avoid operators, loops, conditionals
   + in test code, stick to straight-line code over clever logic, and consider tolerating some duplication when it makes the test more descriptive and meaningful

```java
@Test public void shouldNavigateToPhotosPage() {
  String baseUrl = "http://plus.google.com/";
  Navigator nav = new Navigator(baseUrl);
  nav.goToPhotosPage();
  assertEquals(baseUrl + "/u/0/photos", nav.getCurrentUrl());
}

@Test public void shouldNavigateToPhotosPage() {
  Navigator nav = new Navigator("http://plus.google.com/");
  nav.goToPhotosPage();
  assertEquals("http://plus.google.com//u/0/photos", nav.getCurrentUrl()); // Oops!
}
```

- [**Write clear failure messages**](https://testing.googleblog.com/2014/12/testing-on-toilet-truth-fluent.html)
   + A good failure message contains much the same information as the test's name: it should clearly express the desired outcome, the actual outcome, and any relevant parameters
   + [Truth - Fluent assertions for java and android](https://truth.dev/)

```golang
result := Add(2, 3)
if result != 5 {
    t.Errorf("Add(2, 3) = %v, want %v", result, 5)
}
```

- Tests and code sharing: DAMP(Descriptive and meaningful phrases), not DRY(Don't repeat yourself)
   + [DRY code, DAMP DSLs](http://blog.jayfields.com/2006/05/dry-code-damp-dsls.html)

```java


@Test public void shouldAllowMultipleUsers() {
    User user1 = newUser().setState(State.NORMAL).build();
    User user2 = newUser().setState(State.NORMAL).build();

    Forum forum  = new Forum();
    forum.register(user1);
    forum.register(user2);

    assertThat(forum.hasRegisteredUser(user1)).isTrue();
    assertThat(forum.hasRegisteredUser(user2)).isTrue();
}

@Test public shouldNotRegisterBannedUsers() {
    User user = newUser().setState(State.BANNED).build();
    Forum forum = new Forum();
    try {
        forum.register(user);
    }catch(BannedUserException ignored){}

    assertThat(forum.hasRegisteredUser(user)).isFalse();
}

```
- DAMP is not a replacement for DRY, it is complementary to it.  Helper methods and test infrastructure can still help make tests clearer by making them more concise, factoring out repetitive steps whose details aren't relevant to the particular behavior being tested.

- Shared Values
    + [construct data using helper methods](https://testing.googleblog.com/2018/02/testing-on-toilet-cleanly-create-test.html)
    + Builder pattern: [autovalue](https://github.com/google/auto/tree/master/value)


```java
private static final Account ACCOUNT_1 = Account.newBuilder().setState(AccountState.OPEN).setBalance(50).build();
private static final Account ACCOUNT_2 = Account.newBuilder().setState(AccountState.CLOSED).setBalance(0).build();
private static final ITEM = item.newBuilder().setName("Cheeseburger").setPrice(100).build();
// ....

@Test public void canBuyItem_returnsFalseForClosedAccounts() {
    assertThat(store.canBuyItem(ITEM, ACCOUNT_1)).isFalse();
}

@Test public void canBuyItem_returnsFalseWhenBalanceInsufficient() {
    assertThat(store.canBuyItem(ITEM, ACCOUNT_2)).isFalse();
}

// Change name to CLOSED_ACCOUNT and ACCOUNT_WITH_LOW_BALANCE helps a bit
// difficult to see the exact details of the value being tested

```

```python
def newContact(firstName="Grace", lastName="Hopper", phoneNumber="555-555-5555"):
    return Contact(firstName, lastName, phoneNumber)

# Tests call the helper, specifying values for only the parameters that they care about
def test_fullNameShouldCombineFirstAndLastNames(self):
    def contact = newContact(firstName="Ada", lastName="Lovelace")
    self.assertEqual(contact.fullName(), "Ada Lovelace")

```

```java
// Languages like Java that don't support named parameters can emulate them
// by returning a mutable "builder" object that represents the value under
// construction
private static Contact.Builder newContact() {
    return Contact.newBuilder()
        .setFirstName("Grace")
        .setLastName("Hopper")
        .setPhoneNumber("555-555-5555");
}

// Tests then call methods on the bulder to overwrite only the parameters
// that they care about, then call build() to get a real value out of the
// builder
@Test public void fullNameShouldCombineFirstAndLastNames() {
    Contact contact = newContact()
        .setFirstName("Ada")
        .setLastName("Lovelace")
        .build();
    assertThat(contact.getFullName()).isEqualTo("Ada Lovelace")
}

```

- Shared setup

```java
private NameService nameService;
private UserStore userStore;

@Before public void setUp() {
    nameService = new NameService();
    nameService.set("user1", "Donald Knuth");
    userStore = new UserStore(nameService);
}

// don't use the default value in setUp, try to put all information in test function
@Test public void shouldReturnNameFromService() {
    nameService.set("user1", "Margaret Hamilton");
    UserDetails user = userStore.get("user1");
    assertThat(user.getName()).isEqualTo("Margaret Hamilton");
}

```



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
