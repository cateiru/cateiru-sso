/**********************************************************
 * Documents
 *
 * @author Yoshitsugu Tahara <arisahyper0000@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import Document, {Html, Head, Main, NextScript} from 'next/document';
import {GA_TRACKING_ID} from '../utils/ga/gtag';

export default class MyDocument extends Document {
  render(): JSX.Element {
    return (
      <Html lang="ja">
        <Head>
          {GA_TRACKING_ID && (
            <>
              <script
                async
                src={`https://www.googletagmanager.com/gtag/js?id=${GA_TRACKING_ID}`}
              />
              <script
                dangerouslySetInnerHTML={{
                  __html: `
             window.dataLayer = window.dataLayer || [];
             function gtag(){dataLayer.push(arguments);}
             gtag('js', new Date());
             gtag('config', '${GA_TRACKING_ID}', {
               page_path: window.location.pathname,
             });
               `,
                }}
              />
            </>
          )}
          {/* <link rel="preconnect" href="https://fonts.googleapis.com" />
          <link
            rel="preconnect"
            href="https://fonts.gstatic.com"
            crossOrigin=""
          />
          <link
            rel="apple-touch-icon"
            sizes="180x180"
            href="/favicons/apple-touch-icon.png"
          />
          <link
            rel="icon"
            type="image/png"
            sizes="32x32"
            href="/favicons/favicon-32x32.png"
          />
          <link
            rel="icon"
            type="image/png"
            sizes="16x16"
            href="/favicons/favicon-16x16.png"
          />
          <link rel="manifest" href="/favicons/site.webmanifest" />
          <link
            rel="mask-icon"
            href="/favicons/safari-pinned-tab.svg"
            color="#5bbad5"
          />
          <meta name="msapplication-TileColor" content="#ffffff" />
          <meta
            name="msapplication-config"
            content="/favicons/browserconfig.xml"
          />
          <meta name="theme-color" content="#ffffff"></meta> */}
          {/* ogp */}
          {/* <meta
            name="description"
            content="HelloSlideはあなたと観客がリアルタイムにコミュニケーションを行えるスライドです。"
          />
          <meta property="og:title" content="Hello Slide" />
          <meta
            property="og:description"
            content="HelloSlideはあなたと観客がリアルタイムにコミュニケーションを行えるスライドです。"
          />
          <meta property="og:type" content="website" />
          <meta property="og:url" content="https://hello-slide.jp/" />
          <meta
            property="og:image"
            content={`https://${domain}/images/Card.png`}
          /> */}
          {/* Twitter */}
          {/* <meta name="twitter:card" content="summary_large_image" />
          <meta name="twitter:title" content="Hello Slide" />
          <meta
            name="twitter:description"
            content="HelloSlideはあなたと観客がリアルタイムにコミュニケーションを行えるスライドです。"
          />
          <meta
            name="twitter:image"
            content={`https://${domain}/images/Card.png`}
          />
          <meta name="twitter:site" content="@hello_slide" />
          <meta name="twitter:creator" content="@hello_slide" /> */}
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
