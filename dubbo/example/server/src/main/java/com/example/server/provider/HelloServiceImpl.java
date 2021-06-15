package com.example.server.provider;

import com.alibaba.dubbo.config.annotation.Service;
import com.example.api.HelloService;
import org.springframework.stereotype.Component;

@Service(interfaceClass = HelloService.class)
@Component
public class HelloServiceImpl implements HelloService {
    @Override
    public String SayHello(String userName)  {
        return " Hello "+userName+" from rpc dubbo! ";
    }
}
