#include "b-r-o-w-s-e.h"

@implementation BrowseAppDelegate
- (void)applicationWillFinishLaunching:(NSNotification *)aNotification
{
  NSAppleEventManager *appleEventManager = [NSAppleEventManager sharedAppleEventManager];
  [appleEventManager setEventHandler:self
                         andSelector:@selector(handleGetURLEvent:withReplyEvent:)
                         forEventClass:kInternetEventClass andEventID:kAEGetURL];
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender
{
  return NSTerminateNow;
}

- (void)handleGetURLEvent:(NSAppleEventDescriptor *)event
           withReplyEvent:(NSAppleEventDescriptor *)replyEvent {
  HandleURL((char*)[[[event paramDescriptorForKeyword:keyDirectObject] stringValue] UTF8String]);
}
@end

void RunApp(void) {
  [NSAutoreleasePool new];
  [NSApplication sharedApplication];
  [NSApp setActivationPolicy:NSApplicationActivationPolicyProhibited];
  BrowseAppDelegate *app = [BrowseAppDelegate alloc];
  [NSApp setDelegate:app];
  [NSApp run];
}