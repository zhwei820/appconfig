
# 测试
test:
	go test -v ./services/*


TARGET=$(src)
# 根据模型定义, 生成接口代码
gen_api:$(TARGET)
	python3 gen_api.py $(TARGET)

