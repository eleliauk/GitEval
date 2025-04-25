# 使用 Node.js 20.15.1 作为基础镜像
FROM node:20.15.1

# 创建工作目录
WORKDIR /usr/src/app

# 复制依赖文件
COPY package.json pnpm-lock.yaml* ./

# 安装 pnpm（全局）
RUN npm install -g pnpm

# 安装项目依赖
RUN pnpm install

# 复制项目的其他文件
COPY . .

# 暴露端口
EXPOSE 5173

# 启动应用
CMD ["pnpm", "dev"]
