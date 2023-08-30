import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />
        <link rel="stylesheet" href="node_modules/flexlayout-react/style/dark.css" />
      </Head>
      <body>
        <Main />
        <NextScript />
        <div id="modal_root"></div>
      </body>
    </Html>
  )
}
