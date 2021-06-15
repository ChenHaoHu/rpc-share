package com.example.client.controller;

import com.alibaba.dubbo.config.annotation.Reference;
import com.example.api.HelloService;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1")
public class HelloController {

    @Reference
    HelloService helloService;

    @RequestMapping("/hello")
    public String SayHello(String name){
        name = helloService.SayHello(name);
        //rpc say hello
       return name;
    }
}
