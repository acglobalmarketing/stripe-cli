# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to make participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, sex characteristics, gender identity and expression,
level of experience, education, socio-economic status, nationality, personal
appearance, race, religion, or sexual identity and orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
  advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies within all project spaces, and it also applies when
an individual is representing the project or its community in public spaces.
Examples of representing a project or community include using an official
project e-mail address, posting via an official social media account, or acting
as an appointed representative at an online or offline event. Representation of
a project may be further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at conduct@stripe.com. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at https://www.contributor-covenant.org/version/1/4/code-of-conduct.html

[homepage]: https://www.contributor-covenant.org

For answers to common questions about this code of conduct, see
https://www.contributor-covenant.org/faq
<head>
  <meta charset="utf-8">
  <script src="https://js.braintreegateway.com/web/dropin/1.30.1/js/dropin.min.js"></script>
</head>
<body>
  <!-- Step one: add an empty container to your page -->
  <div id="dropin-container"></div>

  <script type="text/javascript">
  // call `braintree.dropin.create` code here
  </script>
</body>
// Step two: create a dropin instance using that container (or a string
//   that functions as a query selector such as `#dropin-container`)
braintree.dropin.create({
  container: document.getElementById('dropin-container'),
  // ...plus remaining configuration
}, (error, dropinInstance) => {
  // Use `dropinInstance` here
  // Methods documented at https://braintree.github.io/braintree-web-drop-in/docs/current/Dropin.html
});
braintree.dropin.create({
  // Step three: get client token from your server, such as via
 //    templates or async http request
  authorization: CLIENT_TOKEN_FROM_SERVER,
  container: '#dropin-container'
}, (error, dropinInstance) => {
  // Use `dropinInstance` here
  // Methods documented at https://braintree.github.io/braintree-web-drop-in/docs/current/Dropin.html
});
