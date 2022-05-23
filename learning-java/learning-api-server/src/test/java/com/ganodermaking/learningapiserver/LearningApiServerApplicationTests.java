package com.ganodermaking.learningapiserver;

import com.ganodermaking.learningapi.api.IGreeter;
import com.ganodermaking.learningapi.dto.HelloRequest;
import com.ganodermaking.learningapi.dto.HelloResponse;
import org.apache.dubbo.common.utils.Assert;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;

@SpringBootTest
class LearningApiServerApplicationTests {

	@Resource
	private IGreeter greeter;

	@Test
	void contextLoads() {
	}

	@Test
	void hello() {
		HelloRequest request = new HelloRequest();
		request.setName("dubbo");
		HelloResponse response = greeter.sayHello(request);
		Assert.assertTrue(response.getName().startsWith("Hello dubbo"), "fail");
	}
}
