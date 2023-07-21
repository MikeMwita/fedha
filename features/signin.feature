Feature: Signin Feature
  I want to Signin to Fedha API

  Background:
    Given user detail:
      |name               |email              |password  |
      |Jonathan Doe       |johndoe@mail.com   |abcdefg   |
    Given I have load fedha application
    And the fields "email" and "password" are empty

  Scenario: Error on empty fields
    When I click on "Sign In"
    And both fields "email" and "password" have errors
    Then failed to login

  Scenario: Wrong Password
    When I type in "email" and "password"
    And I click on "Sign In"
    Then I should fail to signin and get 'invalid credentials'

  Scenario: Login Successful
    When I type "email" and "password"
    And I click on "Sign In"
    Then the user name should be "Jonathan Doe"
















