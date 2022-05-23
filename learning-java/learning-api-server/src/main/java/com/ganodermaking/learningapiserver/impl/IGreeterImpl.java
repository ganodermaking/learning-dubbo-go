package com.ganodermaking.learningapiserver.impl;

import com.ganodermaking.learningapi.api.IGreeter;
import com.ganodermaking.learningapi.dto.HelloRequest;
import com.ganodermaking.learningapi.dto.HelloResponse;
import org.apache.dubbo.config.annotation.DubboService;

/**
 * @Title: IGreeterImpl.java
 * @Package com.ganodermaking.learningapiserver.facade
 * @Description: (用一句话描述该文件做什么)
 * @Author: 604013
 * @Date: 2022/4/23 3:08 下午
 * @Version V1.0
 * @Copyright: 2022 Shenzhen Hive Box Technology Co.,Ltd All rights reserved.
 * @Note: This content is limited to the internal circulation of Hive Box, and it is prohibited to leak or used for other commercial purposes.
 */
@DubboService
public class IGreeterImpl implements IGreeter {
    @Override
    public HelloResponse sayHello(HelloRequest request) {
        HelloResponse response = new HelloResponse();
        response.setName("Hello " + request.getName());
        return response;
    }
}
