- [How to work in open source community](#how-to-work-in-open-source-community)
  - [Async communication](#async-communication)
  - [Define rules for your project](#define-rules-for-your-project)
  - [How to write commit message](#how-to-write-commit-message)
    - [More information](#more-information)
  - [How to creat a pull request](#how-to-creat-a-pull-request)
  - [Branching rules](#branching-rules)
    - [More information](#more-information-1)
  - [More information](#more-information-2)

# How to work in open source community

## Async communication

Peter Snider recommends this aritical to me [15 rules for communicating at GitHub](https://ben.balter.com/2014/11/06/rules-of-communicating-at-github/), I love it very much.  Some highlights listed below:

1. Prefer asynchronous communication  
2. Don’t underestimate high-fidelity mediums  
Brainstorming  
Feedback  
Small talk and gossip  
3. Nobody gets fired for opening an issue    
Search before you open an issue  
Provide context  
Issues are constructive criticism  
4. Copy teams, not team members  
5. Be mindful of noise  
Avoid drive-by opinions  
Only comment if you have something to add to the discussion  
Always provide context   
6. Use checklists to make blockers explicit  
7. Issues are organization-wide todos  
Issue titles should contain at least one verb  
Don’t close issues prematurely   
8. Master the gentle bump  
9. Keep discussions logically distinct  
discussions should have one purpose, defined by the title at the top of the page.  Discrete topics minimize unnecessary noise and optimize for fast decision making by ensuring only the most relevant teams are involved.  
10. There’s only one way to change something  
11. Secrets, secrets, are no fun  
Email is a terrible, terrible collaboration medium.  It should typically reserved only for things like personnel discussions, one-to-one feedback, and external communication. The same goes for other mediums (like phone calls) that don’t automatically capture and surface context  
12. Surface work early for feedback  
13. If you can’t diff it, don’t use it  
14. Pull requests are community property  
15. Don’t ping, just ask  


## Define rules for your project
Reference come from [robosat#contributing](https://github.com/mapbox/robosat#contributing)
```
- For non-trivial changes you should open a ticket first to outline and discuss ideas and implementation sketches. If you just send us a pull request with thousands of lines of changes we most likely won't accept your changeset.
- We follow the 80/20 rule where 80% of the effects come from 20% of the causes: we strive for simplicity and maintainability over pixel-perfect results. If you can improve the model's accuracy by two percent points but have to add thousands of lines of code we most likely won't accept your changeset.
- We take responsibility for changesets going into master: as soon as your changeset gets approved it is on us to maintain and debug it. If your changeset can not be tested, or maintained in the future by the core developers we most likely won't accept your changeset.
```

## How to write commit message


Recommend [AngularJS Git Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#heading=h.uyo6cb12dt6w)  

### More information
- [How to keep your git commit history clean](https://about.gitlab.com/2018/06/07/keeping-git-commit-history-clean/)
- [阮一峰.Commit message 和 Change log 编写指南](http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html)


## How to creat a pull request
```
#1. List the issue
#2. Tasklist
    - API design
    -  Write failing unit tests for new parameters
    -  Make unit tests pass
    -  Write failing cuke tests
    -  Write code
    -  Review
    -  Adjust for comments
#3. Code Review Checklist - author check these when done, reviewer verify
    -  Code formatted with scripts/format.sh
    -  Changes have test coverage
    -  New exceptions, logging, errors - are messages distinct enough to track down in the code if they get thrown in production on non-debug builds?
    -  The PR is one logically integrated piece of work. If there are unrelated changes, are they at least separate commits?
    -  Commit messages - are they clear enough to understand the intent of the change if we have to look at them later?
    -  Code comments - are there comments explaining the intent?
    -  Relevant docs updated
    -  Changelog entry if required
    -  Impact on the API surface
        -  If HTTP/libosrm.o is backward compatible features, bump the minor version
        -  File format changes require at minor release
        -  If old clients can't use the API after changes, bump the major version

```
Good example could be: [OSRM-3408](https://github.com/Project-OSRM/osrm-backend/pull/3408), [ROCKSDB-3009](https://github.com/facebook/rocksdb/pull/3009), [Turfjs-832](https://github.com/Turfjs/turf/pull/832)


## Branching rules

<object data="https://nvie.com/files/Git-branching-model.pdf" type="application/pdf" width="700px" height="700px">
    <embed src="https://nvie.com/files/Git-branching-model.pdf">
        <p>This browser does not support PDFs. Please download the PDF to view it: <a href="https://nvie.com/files/Git-branching-model.pdf">Download PDF</a>.</p>
    </embed>
</object>

Recommand using [git flow](https://github.com/nvie/gitflow)  

### More information
- [A successful Git branching model](https://nvie.com/posts/a-successful-git-branching-model/)
- [在阿里，我们如何管理代码分支？](https://yq.aliyun.com/articles/573549)
- [git-flow cheatsheet](https://danielkummer.github.io/git-flow-cheatsheet/)


## More information
- [How GitHub Works: Be Asynchronous](https://zachholman.com/posts/how-github-works-asynchronous/)
- [Learn the workings of Git, not just the commands](https://developer.ibm.com/tutorials/d-learn-workings-git/)



