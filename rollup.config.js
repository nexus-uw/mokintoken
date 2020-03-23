import commonjs from '@rollup/plugin-commonjs'
import resolve from '@rollup/plugin-node-resolve'
import builtins from 'rollup-plugin-node-builtins'
import globals from 'rollup-plugin-node-globals'
module.exports = {
  input: 'resources/js/index.js',
  output: {
    dir: 'public'
  },
  plugins: [
    commonjs(),
    resolve({
      browser: true,
      preferBuiltins: false
    }),
    builtins(),
    globals()
  ]
}
