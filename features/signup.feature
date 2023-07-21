Feature:BDD implementation for the signup page

  Background: Load page
    Given a user is on the signup page

  Scenario Outline:A user successfully signs up
    When user enters a "<name>", valid "<email>" and "<password>"
    And user confirms password "<confirm_password>"
    And  user click the signup button
    Then the "<test>" should "<result>" on sign up
    Examples:
      |name               |email              |password  | confirm_password  | test    | result |
      |Jonathan Doe       |johndoe@mail.com   |abcdefg   | abcdefg           | success | pass |
      |Jonathan Doe       |johndoe@mail.com   |abcdefg   | sadvbsadkb           | confirm password | fail |
      |Jonathan Doe       |johndoe   |abcdefg   | abcdefg           | invalid email | fail |


