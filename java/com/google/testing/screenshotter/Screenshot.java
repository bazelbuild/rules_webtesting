// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// //////////////////////////////////////////////////////////////////////////////
//
package com.google.testing.screenshotter;

import com.google.errorprone.annotations.Immutable;
import java.awt.image.BufferedImage;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.util.Base64;
import javax.imageio.ImageIO;
import javax.imageio.ImageReader;
import org.json.JSONObject;

@Immutable
public final class Screenshot {
  private static final String FORMAT = "png";

  private final String base64;

  Screenshot(JSONObject response) {
    this.base64 = response.getString("value");
  }

  public String asBase64() {
    return base64;
  }

  public byte[] asBytes() {
    return Base64.getDecoder().decode(base64);
  }

  public BufferedImage asImage() throws IOException {
    ImageReader imageReader = ImageIO.getImageReadersByFormatName(FORMAT).next();
    imageReader.setInput(ImageIO.createImageInputStream(new ByteArrayInputStream(asBytes())), true);
    return imageReader.read(0);
  }

  @Override
  public boolean equals(Object other) {
    if (other == null) {
      return false;
    }
    if (!(other instanceof Screenshot)) {
      return false;
    }
    Screenshot o = (Screenshot) other;
    return asBase64().equals(o.asBase64());
  }

  @Override
  public int hashCode() {
    return asBase64().hashCode();
  }
}
