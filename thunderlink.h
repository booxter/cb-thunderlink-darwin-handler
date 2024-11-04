#import <Cocoa/Cocoa.h>

extern void HandleURL(char*);

@interface ThunderlinkAppDelegate: NSObject<NSApplicationDelegate>
  - (void)handleGetURLEvent:(NSAppleEventDescriptor *) event withReplyEvent:(NSAppleEventDescriptor *)replyEvent;
@end

void RunApp();
void ShowAlert(char*, char*);
