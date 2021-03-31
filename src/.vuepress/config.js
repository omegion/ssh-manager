const {description} = require('../../package.json')

module.exports = {
    title: 'Bitwarden SSH Manager',
    description: description,

    head: [
        [
            "link",
            {
                rel: "apple-touch-icon",
                sizes: "180x180",
                href: "/favicons/apple-icon-180x180.png",
            },
        ],
        [
            "link",
            {
                rel: "icon",
                type: "image/png",
                sizes: "32x32",
                href: "/favicons/favicon-32x32.png",
            },
        ],
        [
            "link",
            {
                rel: "icon",
                type: "image/png",
                sizes: "16x16",
                href: "/favicons/favicon-16x16.png",
            },
        ],
        ["link", {rel: "shortcut icon", href: "/favicons/favicon.ico"}],
        ["meta", {name: "theme-color", content: "#0842ba"}],
        ["meta", {name: "apple-mobile-web-app-capable", content: "yes"}],
        [
            "meta",
            {name: "apple-mobile-web-app-status-bar-style", content: "black"},
        ],
    ],

    dest: "docs",
    /**
     * Theme configuration, here is the default theme configuration for VuePress.
     *
     * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
     */
    themeConfig: {
        repo: 'https://github.com/omegion/bw-ssh',
        editLinks: false,
        docsDir: '',
        editLinkText: '',
        lastUpdated: false,
        logo: "/img/logo.svg",
        author: "omegion",
        nav: [
            {
                text: 'Guide',
                link: '/guide/',
            },
        ],
        sidebar: {
            '/guide/': [
                {
                    title: 'Guide',
                    collapsable: false,
                    children: [
                        '',
                        'get-started',
                        'quick-start',
                    ]
                }
            ],
        }
    },

    /**
     * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
     */
    plugins: [
        '@vuepress/plugin-back-to-top',
        '@vuepress/plugin-medium-zoom',
    ]
}
