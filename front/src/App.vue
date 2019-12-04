<template>
    <div id="app">
        <div id="nav">
            <el-menu class="el-menu-demo" mode="horizontal">
                <el-menu-item>
                    <router-link to="/">Home</router-link>
                </el-menu-item>
                <el-menu-item>
                    <router-link to="/about">About</router-link>
                </el-menu-item>
                <el-menu-item>
                    <router-link to="/login">Login</router-link>
                </el-menu-item>
                <el-menu-item>
                    <router-link to="/register">Register</router-link>
                </el-menu-item>
                <el-menu-item>
                    <router-link to="/vote">Vote</router-link>
                </el-menu-item>
            </el-menu>
        </div>
        <router-view />
    </div>
</template>

<style lang="scss">
#app {
    color: #2c3e50;
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    text-align: center;
    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;
}

#nav {
    padding-bottom: 30px;

    a {
        color: #2c3e50;
        font-weight: bold;
        text-decoration: none;

        &.router-link-exact-active {
            color: #42b983;
        }
    }
}
</style>
