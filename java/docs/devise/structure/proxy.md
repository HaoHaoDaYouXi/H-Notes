# 代理（Proxy）

## 问题

控制对其它对象的访问。

## 效果

用一个代理来隐藏具体实现类的实现细节，通常还用于在真实的实现的前后添加一部分逻辑。

既然说是代理，那就要对客户端隐藏真实实现，由代理来负责客户端的所有请求。

当然，代理只是个代理，它不会完成实际的业务逻辑，而是一层皮而已，但是对于客户端来说，它必须表现得就是客户端需要的真实实现。

代理有以下四类:
- 远程代理(Remote Proxy): 控制对远程对象(不同地址空间)的访问，它负责将请求及其参数进行编码，并向不同地址空间中的对象发送已经编码的请求。
- 虚拟代理(Virtual Proxy): 根据需要创建开销很大的对象，它可以缓存实体的附加信息，以便延迟对它的访问，例如在网站加载一个很大图片时，不能马上完成，可以用虚拟代理缓存图片的大小信息，然后生成一张临时图片代替原始图片。
- 保护代理(Protection Proxy): 按权限控制对象的访问，它负责检查调用者是否具有实现一个请求所必须的访问权限。
- 智能代理(Smart Reference): 取代了简单的指针，它在访问对象时执行一些附加操作
  - 记录对象的引用次数
  - 当第一次引用一个持久化对象时，将它装入内存
  - 在访问一个实际对象前，检查是否已经锁定了它，以确保其它对象不能改变它。

## 解决方案

代理律师其实就是一种具象化的表现，自己不想回答任何问题，而是委托给律师，律师会根据当事人的请求，回答问题。

```java
public interface People {
    void answerQuestion();
}

public class Parties implements People {
    public boolean answer() {
      // 模拟当事人不想回答问题
      return false;
    }
    
    @Override
    public void answerQuestion() {
      System.out.println("我不知道");
    }
}

public class ProxyLawyer implements People {
    private Parties parties;
  
    public ProxyLawyer(Parties parties) {
      this.parties = parties;
    }
  
    @Override
    public void answerQuestion() {
      while (!parties.answer()) {
        try {
          System.out.println("我当事人不清楚这个");
          Thread.sleep(100);
        } catch (InterruptedException e) {
          e.printStackTrace();
        }
      }
      parties.answerQuestion();
    }
}

public class Client {
    public static void main(String[] args) {
        Parties parties = new Parties();
        ProxyLawyer proxyLawyer = new ProxyLawyer(parties);
        proxyLawyer.answerQuestion();
    }
}
```


----
