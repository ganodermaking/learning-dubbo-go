package com.ganodermaking.learningapi.dto;

import lombok.Data;

import java.io.Serializable;

/**
 * @Title: HelloRequest.java
 * @Package com.ganodermaking.learningapi.dto
 * @Description: (用一句话描述该文件做什么)
 * @Author: 604013
 * @Date: 2022/5/5 3:40 下午
 * @Version V1.0
 * @Copyright: 2022 Shenzhen Hive Box Technology Co.,Ltd All rights reserved.
 * @Note: This content is limited to the internal circulation of Hive Box, and it is prohibited to leak or used for other commercial purposes.
 */
@Data
public class HelloRequest implements Serializable {
    private String name;
}
