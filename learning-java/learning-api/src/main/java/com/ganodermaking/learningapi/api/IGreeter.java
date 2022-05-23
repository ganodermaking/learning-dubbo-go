package com.ganodermaking.learningapi.api;

import com.ganodermaking.learningapi.dto.HelloRequest;
import com.ganodermaking.learningapi.dto.HelloResponse;

/**
 * @author zhouyangzhi
 */
public interface IGreeter {
    /**
     * say hello
     *
     * @param request
     * @return
     */
    HelloResponse sayHello(HelloRequest request);
}
