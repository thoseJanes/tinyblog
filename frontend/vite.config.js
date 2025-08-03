import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    port: 8080,
    // historyApiFallback: {
    //   rewrites: [
    //     { from: /\/index.html\/.*/, to: '/index.html' },
    //     { from: /\/login.html\/.*/, to: '/login.html' }, // 带参数的示例
    //     { from: /\/userpage.html\/.*/, to: '/userpage.html' } // 默认回退
    //   ]
    // },
    proxy: {
      '/api': {  // 所有以 /api 开头的请求
        target: 'http://localhost:3000',  // 后端地址
        changeOrigin: true,
        //rewrite: (path)=>path.replace()
        //pathRewrite: { '^/api': '' }  // 移除 /api 前缀
      }
    }
  },
  resolve: {
    alias: {
      vue: 'vue/dist/vue.esm-bundler.js', // 使用完整版
    },
  },
  pages: {
    mainpage: {
      entry: 'src/main.js',
      template: 'public/index.html',
      filename: 'index.html'
    },
    login: {
      entry: 'src/login.js',
      template: 'public/login.html',
      filename: 'login.html'
    },
    userpage: {
      entry: 'src/userpage.js',
      template: 'public/userpage.html',
      filename: 'userpage.html'
    },
  },
  
  devServer: {
    proxy: {
      '/api': {  // 所有以 /api 开头的请求
        target: 'http://localhost:3000',  // 后端地址
        changeOrigin: true,
        pathRewrite: { '^/api': '' }  // 移除 /api 前缀
      }
    }
  }

})
