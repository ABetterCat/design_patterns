## 1. 定义
```cassandraql
创建一个对象常常需要复杂的过程，所以不适合包含在一个复合对象中。
创建对象可能会导致大量的重复代码，可能会需要复合对象访问不到的信息，
也可能提供不了足够级别的抽象，还可能并不是复合对象概念的一部分。
工厂方法模式通过定义一个单独的创建对象的方法来解决这些问题。
由子类实现这个方法来创建具体类型的对象。
```
## 2.优点
```cassandraql
符合开闭原则，进行了解耦，如果增加类别，使用扩展的方法创建类别，而不是更改创建实例的函数/方法。
他山之石：
1.符合“开闭”原则，具有很强的的扩展性、弹性和可维护性。修改时只需要添加对应的工厂类即可
2.使用了依赖倒置原则，依赖抽象而不是具体，使用（客户）和实现（具体类）松耦合
3.客户只需要知道所需产品的具体工厂，而无须知道具体工厂的创建产品的过程，甚至不需要知道具体产品的类名

链接：https://juejin.im/post/5bdede60518825171a180c44
```
## 3.缺点
```cassandraql
如果不同情况创建不同的类别，还是需要判断，只不过不在工厂创建方法中
他山之石：
1.每增加一个产品时，都需要一个具体类和一个具体创建者，使得类的个数成倍增加，导致系统类数目过多，复杂性增加
2.对简单工厂，增加功能修改的是工厂类；对工厂方法，增加功能修改的是产品类。
```