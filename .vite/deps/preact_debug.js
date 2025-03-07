import "./chunk-76EYESMD.js";
import {
  k,
  l,
  x
} from "./chunk-KT276D35.js";

// node_modules/.deno/preact@10.26.2/node_modules/preact/debug/dist/debug.module.js
var t = {};
function r() {
  t = {};
}
function a(e) {
  return e.type === k ? "Fragment" : "function" == typeof e.type ? e.type.displayName || e.type.name : "string" == typeof e.type ? e.type : "#text";
}
var i = [];
var s = [];
function c() {
  return i.length > 0 ? i[i.length - 1] : null;
}
var l2 = true;
function u(e) {
  return "function" == typeof e.type && e.type != k;
}
function f(n) {
  for (var e = [n], o = n; null != o.__o; ) e.push(o.__o), o = o.__o;
  return e.reduce(function(n2, e2) {
    n2 += "  in " + a(e2);
    var o2 = e2.__source;
    return o2 ? n2 += " (at " + o2.fileName + ":" + o2.lineNumber + ")" : l2 && console.warn("Add @babel/plugin-transform-react-jsx-source to get a more detailed component stack. Note that you should not add it to production builds of your App for bundle size reasons."), l2 = false, n2 + "\n";
  }, "");
}
var d = "function" == typeof WeakMap;
function p(n) {
  var e = [];
  return n.__k ? (n.__k.forEach(function(n2) {
    n2 && "function" == typeof n2.type ? e.push.apply(e, p(n2)) : n2 && "string" == typeof n2.type && e.push(n2.type);
  }), e) : e;
}
function h(n) {
  return n ? "function" == typeof n.type ? null == n.__ ? null != n.__e && null != n.__e.parentNode ? n.__e.parentNode.localName : "" : h(n.__) : n.type : "";
}
var v = x.prototype.setState;
function y(n) {
  return "table" === n || "tfoot" === n || "tbody" === n || "thead" === n || "td" === n || "tr" === n || "th" === n;
}
x.prototype.setState = function(n, e) {
  return null == this.__v && null == this.state && console.warn('Calling "this.setState" inside the constructor of a component is a no-op and might be a bug in your application. Instead, set "this.state = {}" directly.\n\n' + f(c())), v.call(this, n, e);
};
var m = /^(address|article|aside|blockquote|details|div|dl|fieldset|figcaption|figure|footer|form|h1|h2|h3|h4|h5|h6|header|hgroup|hr|main|menu|nav|ol|p|pre|search|section|table|ul)$/;
var b = x.prototype.forceUpdate;
function w(n) {
  var e = n.props, o = a(n), t2 = "";
  for (var r2 in e) if (e.hasOwnProperty(r2) && "children" !== r2) {
    var i2 = e[r2];
    "function" == typeof i2 && (i2 = "function " + (i2.displayName || i2.name) + "() {}"), i2 = Object(i2) !== i2 || i2.toString ? i2 + "" : Object.prototype.toString.call(i2), t2 += " " + r2 + "=" + JSON.stringify(i2);
  }
  var s2 = e.children;
  return "<" + o + t2 + (s2 && s2.length ? ">..</" + o + ">" : " />");
}
x.prototype.forceUpdate = function(n) {
  return null == this.__v ? console.warn('Calling "this.forceUpdate" inside the constructor of a component is a no-op and might be a bug in your application.\n\n' + f(c())) : null == this.__P && console.warn(`Can't call "this.forceUpdate" on an unmounted component. This is a no-op, but it indicates a memory leak in your application. To fix, cancel all subscriptions and asynchronous tasks in the componentWillUnmount method.

` + f(this.__v)), b.call(this, n);
}, l.__m = function(n, e) {
  var o = n.type, t2 = e.map(function(n2) {
    return n2 && n2.localName;
  }).filter(Boolean);
  console.error('Expected a DOM node of type "' + o + '" but found "' + t2.join(", ") + `" as available DOM-node(s), this is caused by the SSR'd HTML containing different DOM-nodes compared to the hydrated one.

` + f(n));
}, function() {
  !function() {
    var n2 = l.__b, o2 = l.diffed, t2 = l.__, r3 = l.vnode, a2 = l.__r;
    l.diffed = function(n3) {
      u(n3) && s.pop(), i.pop(), o2 && o2(n3);
    }, l.__b = function(e) {
      u(e) && i.push(e), n2 && n2(e);
    }, l.__ = function(n3, e) {
      s = [], t2 && t2(n3, e);
    }, l.vnode = function(n3) {
      n3.__o = s.length > 0 ? s[s.length - 1] : null, r3 && r3(n3);
    }, l.__r = function(n3) {
      u(n3) && s.push(n3), a2 && a2(n3);
    };
  }();
  var n = false, o = l.__b, r2 = l.diffed, c2 = l.vnode, l3 = l.__r, v2 = l.__e, b2 = l.__, g = l.__h, E = d ? { useEffect: /* @__PURE__ */ new WeakMap(), useLayoutEffect: /* @__PURE__ */ new WeakMap(), lazyPropTypes: /* @__PURE__ */ new WeakMap() } : null, k2 = [];
  l.__e = function(n2, e, o2, t2) {
    if (e && e.__c && "function" == typeof n2.then) {
      var r3 = n2;
      n2 = new Error("Missing Suspense. The throwing component was: " + a(e));
      for (var i2 = e; i2; i2 = i2.__) if (i2.__c && i2.__c.__c) {
        n2 = r3;
        break;
      }
      if (n2 instanceof Error) throw n2;
    }
    try {
      (t2 = t2 || {}).componentStack = f(e), v2(n2, e, o2, t2), "function" != typeof n2.then && setTimeout(function() {
        throw n2;
      });
    } catch (n3) {
      throw n3;
    }
  }, l.__ = function(n2, e) {
    if (!e) throw new Error("Undefined parent passed to render(), this is the second argument.\nCheck if the element is available in the DOM/has the correct id.");
    var o2;
    switch (e.nodeType) {
      case 1:
      case 11:
      case 9:
        o2 = true;
        break;
      default:
        o2 = false;
    }
    if (!o2) {
      var t2 = a(n2);
      throw new Error("Expected a valid HTML node as a second argument to render.	Received " + e + " instead: render(<" + t2 + " />, " + e + ");");
    }
    b2 && b2(n2, e);
  }, l.__b = function(e) {
    var r3 = e.type;
    if (n = true, void 0 === r3) throw new Error("Undefined component passed to createElement()\n\nYou likely forgot to export your component or might have mixed up default and named imports" + w(e) + "\n\n" + f(e));
    if (null != r3 && "object" == typeof r3) {
      if (void 0 !== r3.__k && void 0 !== r3.__e) throw new Error("Invalid type passed to createElement(): " + r3 + "\n\nDid you accidentally pass a JSX literal as JSX twice?\n\n  let My" + a(e) + " = " + w(r3) + ";\n  let vnode = <My" + a(e) + " />;\n\nThis usually happens when you export a JSX literal and not the component.\n\n" + f(e));
      throw new Error("Invalid type passed to createElement(): " + (Array.isArray(r3) ? "array" : r3));
    }
    if (void 0 !== e.ref && "function" != typeof e.ref && "object" != typeof e.ref && !("$$typeof" in e)) throw new Error(`Component's "ref" property should be a function, or an object created by createRef(), but got [` + typeof e.ref + "] instead\n" + w(e) + "\n\n" + f(e));
    if ("string" == typeof e.type) {
      for (var i2 in e.props) if ("o" === i2[0] && "n" === i2[1] && "function" != typeof e.props[i2] && null != e.props[i2]) throw new Error(`Component's "` + i2 + '" property should be a function, but got [' + typeof e.props[i2] + "] instead\n" + w(e) + "\n\n" + f(e));
    }
    if ("function" == typeof e.type && e.type.propTypes) {
      if ("Lazy" === e.type.displayName && E && !E.lazyPropTypes.has(e.type)) {
        var s2 = "PropTypes are not supported on lazy(). Use propTypes on the wrapped component itself. ";
        try {
          var c3 = e.type();
          E.lazyPropTypes.set(e.type, true), console.warn(s2 + "Component wrapped in lazy() is " + a(c3));
        } catch (n2) {
          console.warn(s2 + "We will log the wrapped component's name once it is loaded.");
        }
      }
      var l4 = e.props;
      e.type.__f && delete (l4 = function(n2, e2) {
        for (var o2 in e2) n2[o2] = e2[o2];
        return n2;
      }({}, l4)).ref, function(n2, e2, o2, r4, a2) {
        Object.keys(n2).forEach(function(o3) {
          var i3;
          try {
            i3 = n2[o3](e2, o3, r4, "prop", null, "SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED");
          } catch (n3) {
            i3 = n3;
          }
          i3 && !(i3.message in t) && (t[i3.message] = true, console.error("Failed prop type: " + i3.message + (a2 && "\n" + a2() || "")));
        });
      }(e.type.propTypes, l4, 0, a(e), function() {
        return f(e);
      });
    }
    o && o(e);
  };
  var T, _ = 0;
  l.__r = function(e) {
    l3 && l3(e), n = true;
    var o2 = e.__c;
    if (o2 === T ? _++ : _ = 1, _ >= 25) throw new Error("Too many re-renders. This is limited to prevent an infinite loop which may lock up your browser. The component causing this is: " + a(e));
    T = o2;
  }, l.__h = function(e, o2, t2) {
    if (!e || !n) throw new Error("Hook can only be invoked from render methods.");
    g && g(e, o2, t2);
  };
  var O = function(n2, e) {
    return { get: function() {
      var o2 = "get" + n2 + e;
      k2 && k2.indexOf(o2) < 0 && (k2.push(o2), console.warn("getting vnode." + n2 + " is deprecated, " + e));
    }, set: function() {
      var o2 = "set" + n2 + e;
      k2 && k2.indexOf(o2) < 0 && (k2.push(o2), console.warn("setting vnode." + n2 + " is not allowed, " + e));
    } };
  }, I = { nodeName: O("nodeName", "use vnode.type"), attributes: O("attributes", "use vnode.props"), children: O("children", "use vnode.props.children") }, M = Object.create({}, I);
  l.vnode = function(n2) {
    var e = n2.props;
    if (null !== n2.type && null != e && ("__source" in e || "__self" in e)) {
      var o2 = n2.props = {};
      for (var t2 in e) {
        var r3 = e[t2];
        "__source" === t2 ? n2.__source = r3 : "__self" === t2 ? n2.__self = r3 : o2[t2] = r3;
      }
    }
    n2.__proto__ = M, c2 && c2(n2);
  }, l.diffed = function(e) {
    var o2, t2 = e.type, i2 = e.__;
    if (e.__k && e.__k.forEach(function(n2) {
      if ("object" == typeof n2 && n2 && void 0 === n2.type) {
        var o3 = Object.keys(n2).join(",");
        throw new Error("Objects are not valid as a child. Encountered an object with the keys {" + o3 + "}.\n\n" + f(e));
      }
    }), e.__c === T && (_ = 0), "string" == typeof t2 && (y(t2) || "p" === t2 || "a" === t2 || "button" === t2)) {
      var s2 = h(i2);
      if ("" !== s2 && y(t2)) "table" === t2 && "td" !== s2 && y(s2) ? (console.log(s2, i2.__e), console.error("Improper nesting of table. Your <table> should not have a table-node parent." + w(e) + "\n\n" + f(e))) : "thead" !== t2 && "tfoot" !== t2 && "tbody" !== t2 || "table" === s2 ? "tr" === t2 && "thead" !== s2 && "tfoot" !== s2 && "tbody" !== s2 ? console.error("Improper nesting of table. Your <tr> should have a <thead/tbody/tfoot> parent." + w(e) + "\n\n" + f(e)) : "td" === t2 && "tr" !== s2 ? console.error("Improper nesting of table. Your <td> should have a <tr> parent." + w(e) + "\n\n" + f(e)) : "th" === t2 && "tr" !== s2 && console.error("Improper nesting of table. Your <th> should have a <tr>." + w(e) + "\n\n" + f(e)) : console.error("Improper nesting of table. Your <thead/tbody/tfoot> should have a <table> parent." + w(e) + "\n\n" + f(e));
      else if ("p" === t2) {
        var c3 = p(e).filter(function(n2) {
          return m.test(n2);
        });
        c3.length && console.error("Improper nesting of paragraph. Your <p> should not have " + c3.join(", ") + " as child-elements." + w(e) + "\n\n" + f(e));
      } else "a" !== t2 && "button" !== t2 || -1 !== p(e).indexOf(t2) && console.error("Improper nesting of interactive content. Your <" + t2 + "> should not have other " + ("a" === t2 ? "anchor" : "button") + " tags as child-elements." + w(e) + "\n\n" + f(e));
    }
    if (n = false, r2 && r2(e), null != e.__k) for (var l4 = [], u2 = 0; u2 < e.__k.length; u2++) {
      var d2 = e.__k[u2];
      if (d2 && null != d2.key) {
        var v3 = d2.key;
        if (-1 !== l4.indexOf(v3)) {
          console.error('Following component has two or more children with the same key attribute: "' + v3 + '". This may cause glitches and misbehavior in rendering process. Component: \n\n' + w(e) + "\n\n" + f(e));
          break;
        }
        l4.push(v3);
      }
    }
    if (null != e.__c && null != e.__c.__H) {
      var b3 = e.__c.__H.__;
      if (b3) for (var g2 = 0; g2 < b3.length; g2 += 1) {
        var E2 = b3[g2];
        if (E2.__H) {
          for (var k3 = 0; k3 < E2.__H.length; k3++) if ((o2 = E2.__H[k3]) != o2) {
            var O2 = a(e);
            console.warn("Invalid argument passed to hook. Hooks should not be called with NaN in the dependency array. Hook index " + g2 + " in component " + O2 + " was called with NaN.");
          }
        }
      }
    }
  };
}();
export {
  r as resetPropWarnings
};
//# sourceMappingURL=preact_debug.js.map
