import { useState } from 'react'
import './App.less'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div className="card animate-pulse">
        <a className={`repo`} href={`https://github.com/LightXi/allproxy`}>
          <svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title>GitHub</title><path d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"></path></svg>
        </a>
        <div className={`card-header`}>
          <img src={`/logo.png`} alt="" />
          <span>LightXi 开放平台</span>
        </div>
        <div className={`card-body`}>
          <div className={`app`}>
            <a className={`title`} href={`https://googlefonts.cn`} target={`_blank`}>Google Fonts</a>
            <div className={`code`}>
              https://open.lightxi.com/fonts/<strong className={`point`}>Jetbrains-Mono</strong>
            </div>
          </div>

          <div className={`app`}>
            <a className={`title`} href={`https://jsdelivr.com`} target={`_blank`}>JsDelivr</a>
            <div className={`code`}>
              https://open.lightxi.com/jsdelivr/<strong className={`point`}>npm/jquery@3.7.1/dist/jquery.min.js</strong>
            </div>
          </div>

          <div className={`app`}>
            <a className={`title`} href={`https://esm.sh`} target={`_blank`}>ESM</a>
            <div className={`code`}>
              https://open.lightxi.com/esm/<strong className={`point`}>canvas-confetti@1.6.1</strong>
            </div>
          </div>

          <div className={`app`}>
            <a className={`title`} href={`https://unpkg.com`} target={`_blank`}>Unpkg</a>
            <div className={`code`}>
              https://open.lightxi.com/unpkg/<strong className={`point`}>react@16.7.0/umd/react.production.min.js</strong>
            </div>
          </div>
        </div>
        <div className={`card-footer`}>
          powered by
          <a href={`https://lightxi.com`} target={`_blank`}>
            <img src={`/logo.png`} alt="" />
            LightXi
          </a>
        </div>
      </div>
      <div className={`footer`}>
        <a href="https://beian.miit.gov.cn/" target="_blank">
          <img src="https://lightxi.com/themes/web/mfQloud/assets/images/gov.webp" alt="" />
            粤ICP备2023066011号-1
        </a>
        <a href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=44040302000432" target="_blank">
          <img src="https://lightxi.com/themes/web/mfQloud/assets/images/gongan.webp" alt="" />
            粤公网安备 44040302000432号
        </a>
        <a href="https://dxzhgl.miit.gov.cn/dxxzsp/xkz/xkzgl/resource/qiyesearch.jsp?num=B1-20234898&amp;type=xuke" target="_blank">
          <img src="https://lightxi.com/themes/web/mfQloud/assets/images/cert.webp" alt="" />
            增值电信业务经营许可证 B1-20234898
        </a>
      </div>
    </>
  )
}

export default App
