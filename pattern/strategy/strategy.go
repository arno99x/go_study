/**
其实本例是策略模式+简单工厂模式

策略模式：要求调用者对业务比较了解，由创建者直接创建好类让策略模式去执行即可
优点：
	1. 由于各产品实现同样的interface，所以可以互相切换调用，让客户端可以根据业务需要去调用不同的业务
	2. 添加新的产品，不需策略模式修改代码，调用方修改就行。 （但结合了简单工厂模式后，简单工厂是要修改一下的）
缺点:
	1. 调用者必须清晰的知道每个策略类的功能
	2. 策略类多了，维护起来会特别麻烦
适用场景
	做面向对象设计的，对策略模式一定很熟悉，因为它实质上就是面向对象中的继承和多态，在看完策略模式的通用代码后，
	我想，即使之前从来没有听说过策略模式，在开发过程中也一定使用过它吧？至少在在以下两种情况下，大家可以考虑使
	用策略模式，几个类的主要逻辑相同，只在部分逻辑的算法和行为上稍有区别的情况。有几种相似的行为，或者说算法，
	客户端需要动态地决定使用哪一种，那么可以使用策略模式，将这些算法封装起来供客户端调用。
	策略模式是一种简单常用的模式，我们在进行开发的时候，会经常有意无意地使用它，一般来说，策略模式不会单独使用，
	跟模版方法模式、工厂模式等混合使用的情况比较多。
*/
package strategy

type CashSuper interface {
	AcceptCash(money float32) float32
}

type CashNomal struct {
}

type CashRebate struct {
	meneyRebate float32
}

type CashReturn struct {
	meneyCondition float32
	meneyReturn    float32
}

func (CashNomal) AcceptCash(money float32) float32 {

	return money
}

func (c *CashRebate) AcceptCash(money float32) float32 {
	return c.meneyRebate * money
}

func (c *CashReturn) AcceptCash(money float32) float32 {
	if money >= c.meneyCondition {
		return money - float32(int(money/c.meneyCondition*c.meneyReturn))
	} else {
		return money
	}
}

type CashContext struct {
	CashSuper
}

func NewCashContext(context string) CashSuper {
	switch context {
	case "300减100":
		return &CashReturn{300, 100}
	case "normal":
		return CashNomal{}
	case "打八折":
		return &CashRebate{0.8}
	default:
		return nil
	}
}
