# 拉取官方镜像
FROM postgres:11.4
# 设置作者名字（选配）
LABEL author="lhs"

# 设置时区（建议配置）
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

#初始化PostgreSQL
# super user 用户名（建议配置）
ENV POSTGRES_USER root

# super user 密码（建议配置）
ENV POSTGRES_PASSWORD Abc123++

# 创建一个数据库，设置其名称（建议配置）
ENV POSTGRES_DB db

# 设置docker中的数据存储路径（选配）
# 默认值为/var/lib/postgresql/data，手动配置后会将数据文件存储到新的位置
ENV PGDATA=/var/lib/postgresql/data/pgdata

#容器运行时监听的端口
EXPOSE 5432
