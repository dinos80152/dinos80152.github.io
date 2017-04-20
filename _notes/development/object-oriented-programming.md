# Object Oriented Programming

## Basic

### Visibility

|           | Object | SubClass | SelfClass |
| --------- | ------ | -------- | --------- |
| public    | v      | v        | v         |
| protected | x      | v        | v         |
| private   | x      | x        | v         |

### Class

* Constructor
* Field
  * variable
  * constant
* Method
  * Regular
    * depends on object
  * Static
    * depends on class
* Destructor

### Encapsulation

don't operate data directly, use methods by *parameters* and *arguments*:

* Getter, Accessors
* Setter, Mutator

### Inherited

* IS-A
  * Is a *Dog* an *Animal*?
* HAS-A
  * *Dog* has a *Height*.

### Polymorphism

SubClass can override SuperClass methods, and be used when it's declared as a SuperClass object.

#### Abstract Class

* It's like "theory of Forms"(Plato) in philosophy. - Dino Lai
* use the power of Polymorphism without the work.
* all the method in abstract should be abstract.
* no abstract fields.

### Interface

* It's like contract with employee and employer. - Dino Lai
* Only declares methods to be implemented.