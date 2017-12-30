# Authorization Models: ACL, DAC, MAC, RBAC, ABAC

## ACL (Access Control List)

* `Subject` can `Action` to `Object`
* Base on user and group

### Example

1. Granting Dino article created permission.

    ```yaml
    Subject: Dino
    Action: Create
    Object: Article
    ```

1. Dino can create article now.

## DAC (Discretionary Access Control)

> The controls are discretionary in the sense that a subject with a certain access permission is capable of passing that permission (perhaps indirectly) on to any other subject.

* `Subject` can `Action` to `Object`
* `Subject` can `grant` other `Subject`
* Base on user and group

### Example

1. Granting Dino article created permission.

    ```yaml
    Subject: Dino
    Action: Create
    Object: Article
    ```

1. Dino can create article now, and give this permission to others.

1. Dino grants James to create articles.

    ```yaml
    Subject: James
    Action: Create
    Object: Article
    ```

1. James can create article now.

## MAC (Mandatory Access Control)

> **Subjects and objects each have a set of security attributes.** Whenever a subject attempts to access an object, an authorization rule enforced by the operating system kernel examines these security attributes and decides whether the access can take place. Any operation by any subject on any object is tested against the set of authorization rules (aka policy) to determine if the operation is allowed.

> With mandatory access control, this security policy is centrally controlled by a security policy administrator; users do not have the ability to override the policy and, for example, grant access to files that would otherwise be restricted.

* `Subject` can `Action` to `Object`
* `Object` can be `Action` by `Subject`
* Base on user and group

### Example

1. Granting Dino article created permission.

    ```yaml
    Subject: Dino
    Action: Create
    Object: Article
    ```

1. Let Article could be created by Dino.

    ```yaml
    Subject: Article
    Action: Created
    Object: Dino
    ```

1. Dino can create article now.

## RBAC (Role-Based Access Control)

> RBAC differs from access control lists (ACLs), used in traditional discretionary access-control systems, in that **it assigns permissions to specific operations with meaning in the organization**, rather than to low level data objects. For example, an access control list could be used to grant or deny write access to a particular system file, but it would not dictate how that file could be changed. In an RBAC-based system, an operation might be to 'create a credit account' transaction in a financial application or to 'populate a blood sugar level test' record in a medical application.

* `Subject` is a `Role` which has `Permission` of `Action` to `Object`
* Can implement mandatory access control (MAC) or discretionary access control (DAC).
* (User or group)-role-permission-object
* Concept
    * Subject
    * Role
    * Permission
    * Operation

### Group vs Role

* Group: a collection of users
    * Dino, James and Liam are members of Meifamly Organization.
* Role: a collection of permissions
    * Writer is a role, which can create, update articles.
    * Role can be applied to user and group.

### Example

1. Set permissions named `write article` and `manage article`

    ```yaml
    Permission:
        - Name: write article
        - Operations:
            - Object: Article
              Action: Created
            - Object: Article
              Action: Updated
            - Object: Article
              Action: Read
        - Name: manage article
        - Operations:
            - Object: Article
              Action: Delete
            - Object: Article
              Action: Read
    ```

1. Set a Role named `Writer`, give it `write article` permission, and a Role named `Manager`, give it `manage article` permission. CEO has all permissions.

    ```yaml
    Role:
        - Name: Writer
          Permissions:
            - write article
        - Name: Manager
          Permissions:
            - manage article
        - Name: CEO
          Permissions:
            - write article
            - manage article
    ```

1. Give Dino `Writer` role

    ```yaml
    Subject: Dino
    Roles:
        - Writer
    ```

1. Dino can create article now.

1. Give James `Writer` and `Manager` roles

    ```yaml
    Subject: James
    Roles:
        - Writer
        - Manager
    ```

1. James can create and delete article now.

## ABAC (Attribute-Based Access Control)

> Unlike role-based access control (RBAC), which employs pre-defined roles that carry a specific set of privileges associated with them and to which subjects are assigned, **the key difference with ABAC is the concept of policies that express a complex Boolean rule set that can evaluate many different attributes.**

* `Subject` who is xxx can `Action` to `Object` which is xxx in `Environment`.
* Concept
    * Policies: bring together attributes to express what can happen and is not allowed.
    * Attributes
        * Subject
            * age, clearance, department, role, job title.
        * Action
            * read, delete, view, approve
        * Resource
            * the object type (medical record, bank account...), the department, the classification or sensitivity, the location
        * Contextual (environment)
            * attributes that deal with time, location or dynamic aspects of the access control scenario
* Standard
    * [XACML](https://en.wikipedia.org/wiki/XACML) (eXtensible Access Control Markup Language)

### Example

* Dino who in Product Department as a Writer could create and update the article, which tag is technology and software in draft mode, and the connection is from Taiwan between 2017-12-01 and 2017-12-31.

    ```yaml
    Subject:
        Name: Dino
        Department: Product
        Role: Writer
    Action:
        - create
        - update
    Resource:
        Type: Article
        Tag:
            - technology
            - software
        Mode:
            - draft
    Contextual:
        Location: Taiwan
        StartTime: 2017-12-01
        EndTime: 2017-12-31
    ```

### AWS Resource-Based Policies is a kind of ABAC

* Limits Terminating EC2 Instances to an IP Address Range

    ```json
    {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:TerminateInstances"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Effect": "Deny",
            "Action": [
                "ec2:TerminateInstances"
            ],
            "Condition": {"NotIpAddress": {"aws:SourceIp": [
                "192.0.2.0/24",
                "203.0.113.0/24"
                ]}},
            "Resource": [
                "arn:aws:ec2:<REGION>:<ACCOUNTNUMBER>:instance/*"
            ]
        }
    ]
    }
    ```

## Reference

* [Access control list@Wiki](https://en.wikipedia.org/wiki/Access_control_list)
* [Discretionary access control@wiki](https://en.wikipedia.org/wiki/Discretionary_access_control)
* [Mandatory access control@wiki](https://en.wikipedia.org/wiki/Mandatory_access_control)
* [role-based access control@wiki](https://en.wikipedia.org/wiki/Role-based_access_control)
* [Attribute-based access control@wiki](https://en.wikipedia.org/wiki/Attribute-based_access_control)
* [XACML@wiki](https://en.wikipedia.org/wiki/XACML)
* [Group vs Role@Stackoverflow](https://stackoverflow.com/questions/7770728/group-vs-role-any-real-difference)
* [IAM Policies Example@AWS User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_ec2_terminate-ip.html)