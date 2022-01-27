# 创建账户

url: /api/v1/createAccount

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| name | 用户名 | String | 是 |
| password | 密码 | String | 是 |
| mnemonic | 助记词 | String | 否 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| address | 地址 | String |
| mnemonic | 助记词 | String |

### 请求示例
```json
{
    "name": "lghh",
    "password": "Qwer1234"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "address": "iaa1uv5pqqqt5p2am2kstvwylyvfgn0glzsuy73dpa",
    "mnemonic": "script service horn crop hood over link scale bubble sword right portion pistol trophy luxury employ toddler east verify unfold camp kick hamster wheat"
}
```

# 查询账户

url: /api/v1/queryAccount

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| name | 用户名 | String | 是 |
| password | 密码 | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| address | 地址 | String |

### 请求示例
```json
{
    "name": "lghh",
    "password": "Qwer1234"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "address": "iaa1uv5pqqqt5p2am2kstvwylyvfgn0glzsuy73dpa"
}
```

# 创建NFT类别

url: /api/v1/issueDenom

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| id | ID | String | 是 |
| name | 名字 | String | 是 |
| schema | 类别信息 | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| gas_wanted | 投入的gas | Integer |
| gas_used | 消耗的gas | Integer |
| hash | 交易哈希 | String |
| height | 区块高度 | Integer |

### 请求示例
```json
{
    "id": "lghdenom",
    "name": "lghdenom",
    "schema": "{}"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "gas_wanted": 200000,
    "gas_used": 68487,
    "hash": "829F7E56E631581B810F5ED7E29817E77570FB99C7D3491B2898D7878587F952",
    "height": 2929448
}
```

# 创建NFT

url: /api/v1/mintNFT

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| id | ID | String | 是 |
| name | 名字 | String | 是 |
| uri | nft的uri | String | 是 |
| data | nft数据 | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| gas_wanted | 投入的gas | Integer |
| gas_used | 消耗的gas | Integer |
| hash | 交易哈希 | String |
| height | 区块高度 | Integer |

### 请求示例
```json
{
    "denomId": "lghdenom",
    "id": "lghnft",
    "name": "lghnft",
    "uri": "http://...",
    "data": "lgh"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "gas_wanted": 200000,
    "gas_used": 73407,
    "hash": "8885FC4B14D1FDC8F2251414A3CCD7C9F2CD38A4022C707427F49BF8A8FFD22F",
    "height": 2929489
}
```

# 转移NFT

url: /api/v1/transferNFT

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| denomId | NFT类型ID | String | 是 |
| id | NFT的ID | String | 是 |
| recipient | 接收者地址 | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| gas_wanted | 投入的gas | Integer |
| gas_used | 消耗的gas | Integer |
| hash | 交易哈希 | String |
| height | 区块高度 | Integer |

### 请求示例
```json
{
    "denomId": "lghdenom",
    "id": "lghnft",
    "recipient": "iaa1uv5pqqqt5p2am2kstvwylyvfgn0glzsuy73dpa"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "gas_wanted": 200000,
    "gas_used": 70800,
    "hash": "92E0BF0A3B1AD60D96B0D9A59EF7001740D46279C9515943C21E626813FFFEC9",
    "height": 2929515
}
```

# 查询NFT类别

url: /api/v1/queryDenom

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| id | ID | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| id | ID | String |
| name | 名字 | String |
| schema | 类别信息 | String |
| creator | 创建者 | String |

### 请求示例
```json
{
    "id": "lghdenom"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "id": "lghdenom",
    "name": "lghdenom",
    "schema": "{}",
    "creator": "iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f"
}
```

# 查询NFT

url: /api/v1/queryNFT

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| id | ID | String | 是 |
| denomId | NFT类别ID | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| id | ID | String |
| name | 名字 | String |
| uri | nft的uri | String |
| data | nft数据 | String |
| creator | 创建者 | String |

### 请求示例
```json
{
    "denomId": "lghdenom",
    "id": "lghnft"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "id": "lghnft",
    "name": "lghnft",
    "uri": "http://...",
    "data": "lgh",
    "creator": "iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f"
}
```

# 查询NFT列表

url: /api/v1/queryNFTs

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| denomId | NFT类别ID | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| id | ID | String |
| name | 名字 | String |
| schema | 类别信息 | String |
| creator | 创建者 | String |
| nfts | NFT列表 | Array of NFT |

### NFT结构
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| id | ID | String |
| name | 名字 | String |
| uri | nft的uri | String |
| data | nft数据 | String |
| creator | 创建者 | String |

### 请求示例
```json
{
    "denomId": "lghdenom"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "id": "lghdenom",
    "name": "lghdenom",
    "schema": "{}",
    "creator": "iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f",
    "nfts": [
        {
            "id": "lghnft",
            "name": "lghnft",
            "uri": "http://...",
            "data": "lgh",
            "creator": "iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f"
        }
    ]
}
```

# 查询拥有的NFT列表

url: /api/v1/queryOwnNFTs

method: POST

### 请求参数
| 参数名 | 参数描述 | 参数类型 | 是否必填 |
|:----------:|-------------|-------------|-------------|
| address | 地址 | String | 是 |
| denomId | NFT类别ID | String | 是 |

### 返回参数
| 参数名 | 参数描述 | 参数类型 |
|:----------:|-------------|-------------|
| code | 错误码 | Integer |
| message | 错误信息 | String |
| ids | NFT的ID列表 | Array of String |

### 请求示例
```json
{
    "address": "iaa1uv5pqqqt5p2am2kstvwylyvfgn0glzsuy73dpa",
    "denomId": "lghdenom"
}
```

### 返回示例
```json
{
    "code": 0,
    "message": "",
    "ids": [
        "lghnft"
    ]
}
```

# 部署流程

安装golang 1.14.15

执行以下命令编译

```
make irita-api
```

根据实际情况修改bin/config/irita_config.yaml中的配置

```yaml
server:
  port: 8080

logger:
  level: debug

irita:
  username: test_key_name
  password: test_password
  mnemonic: supreme zero ladder chaos blur lake dinner warm rely voyage scan dilemma future spin victory glance legend faculty join man mansion water mansion exotic
  coin: 100000uirita
  gasLimit: 200000
  tlsEnable: false
  rpcAddress: http://47.100.192.234:26657
  wsAddress: ws://47.100.192.234:26657
  grpcAddress: 47.100.192.234:9090
  chainId: testing
  projectId: TestProjectID
  projectKey: TestProjectKey
  chainAccountAddress: TestChainAccountAddress
```

在`bin`目录下执行以下命令启动
```
./irita-api
```