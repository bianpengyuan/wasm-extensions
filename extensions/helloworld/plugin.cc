#include <string>
#include <string_view>
#include <unordered_map>

#include "proxy_wasm_intrinsics.h"

class HelloWorldRootContext : public RootContext {
 public:
  explicit HelloWorldRootContext(uint32_t id, std::string_view root_id)
      : RootContext(id, root_id) {}
};

class HelloWorldContext : public Context {
 public:
  explicit HelloWorldContext(uint32_t id, RootContext* root)
      : Context(id, root) {}

  FilterHeadersStatus onResponseHeaders(uint32_t headers,
                                        bool end_of_stream) override;
};

FilterHeadersStatus HelloWorldContext::onResponseHeaders(uint32_t, bool) {
  addResponseHeader("hello", "world");
  return FilterHeadersStatus::Continue;
}

static RegisterContextFactory register_HelloWorldContext(
    CONTEXT_FACTORY(HelloWorldContext), ROOT_FACTORY(HelloWorldRootContext));
